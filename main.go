package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"unsafe"

	tea "github.com/charmbracelet/bubbletea"
)

var factions [4]string

type Planet struct {
	Name          string
	Hash          int
	Index         int
	Initial_owner string
	Max_health    int
}

type SinglePlanet struct {
	Health           int
	Liberation       int
	Players          int
	Regen_per_second int
}

type State struct {
	AllPlanets       []Planet
	SelectedIdx      int
	FactionFilterIdx int
	ActivePlanet     SinglePlanet
}

type AllPlanetsMsg []Planet

type SinglePlanetMsg SinglePlanet

func initialModel() State {
	return State{
		AllPlanets:       make([]Planet, 0),
		SelectedIdx:      260,
		FactionFilterIdx: 0,
		ActivePlanet:     SinglePlanet{},
	}
}

func (m State) Init() tea.Cmd {
	return fetchAllPlanets
}

func (m State) View() string {

	activeFaction := factions[m.FactionFilterIdx%4]
	// The header
	s := fmt.Sprintf("Planets (%s)\n\n", activeFaction)

	planetsToRender := m.AllPlanets
	if activeFaction != "All" {
		planetsToRender = make([]Planet, 0)
		for _, planet := range m.AllPlanets {
			if planet.Initial_owner == activeFaction {
				planetsToRender = append(planetsToRender, planet)
			}
		}

	}

	// Iterate over our choices
	for idx, planet := range planetsToRender {

		selected := ""
		if idx == m.SelectedIdx {
			selected = "X"
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", planet.Name, selected)
	}

	if unsafe.Sizeof(m.ActivePlanet) != 0 {
		s += fmt.Sprintf(" \n\n%d %% liberated ", m.ActivePlanet.Liberation)
	}

	return s
}

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case AllPlanetsMsg:
		m.AllPlanets = msg

	case SinglePlanetMsg:

		m.ActivePlanet = SinglePlanet{
			Health:           msg.Health,
			Liberation:       msg.Liberation,
			Players:          msg.Players,
			Regen_per_second: msg.Regen_per_second,
		}

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.SelectedIdx > 0 {
				m.SelectedIdx--
			}
		case "down":
			if m.SelectedIdx < len(m.AllPlanets)-1 {
				m.SelectedIdx++
			}

		case "enter":

			return m, fetchPlanetInfo(m.AllPlanets[m.SelectedIdx].Index)

		case "tab":
			m.FactionFilterIdx++
			m.SelectedIdx = 0
		}

	}
	return m, nil
}

func main() {

	factions = [...]string{"All", "Humans", "Terminids", "Automaton"}

	program := tea.NewProgram(initialModel())
	if _, err := program.Run(); err != nil {
		os.Exit(1)
	}
}

func fetchAllPlanets() tea.Msg {

	client := &http.Client{Timeout: 10 * time.Second}
	res := makeHttpRequest[[]Planet](client, "https://helldivers-2.fly.dev/api/801/planets")
	return AllPlanetsMsg(res)
}

func fetchPlanetInfo(id int) tea.Cmd {

	return func() tea.Msg {
		client := &http.Client{Timeout: 10 * time.Second}
		res := makeHttpRequest[SinglePlanet](client, fmt.Sprintf("https://helldivers-2.fly.dev/api/801/planets/%d/status", 196))

		return SinglePlanetMsg(res)
	}

}

func makeHttpRequest[T any](httpClient *http.Client, url string) T {
	res, err := httpClient.Get(url)

	if err != nil {

		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	var AllPlanets T

	json.Unmarshal(body, &AllPlanets)

	return AllPlanets

}
