package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"unsafe"

	tea "github.com/charmbracelet/bubbletea"
)

var client = &http.Client{Timeout: 10 * time.Second}
var factions [4]string

func (m State) Init() tea.Cmd {
	return fetchAllPlanets_msg
}

func (m State) View() string {
	activeFaction := factions[m.FactionFilterIdx%4]

	// The header
	s := fmt.Sprintf("Planets (%s)\n", activeFaction)

	// Iterate over our choices
	for idx, planet := range m.FilteredPlanets {

		selected := " "
		if idx == m.SelectedIdx {
			selected = ">"
		}

		text := fmt.Sprintf("%s %s\n", selected, planet.Name)

		switch planet.Initial_owner {

		case "Humans":
			s += HumanText.Render(text)
		case "Terminids":
			s += TerminidText.Render(text)
		case "Automaton":
			s += AutomatonText.Render(text)
		}

	}

	if unsafe.Sizeof(m.ActivePlanet) != 0 {
		s += fmt.Sprintf(" \n\n%f %% liberated ", m.ActivePlanet.Liberation)
	}

	return s
}

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case AllPlanetsMsg:
		m.AllPlanets = msg

	case SinglePlanetMsg:
		m.ActivePlanet = SinglePlanet(msg)

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.SelectedIdx > 0 {
				m.SelectedIdx--
			}
		case "down":
			if m.SelectedIdx < len(m.FilteredPlanets)-1 {
				m.SelectedIdx++
			}

		case "enter":

			return m, fetchPlanetInfo(m.FilteredPlanets[m.SelectedIdx].Index)

		case "tab":
			m.FactionFilterIdx++

			activeFaction := factions[m.FactionFilterIdx%4]
			activePlanets := m.AllPlanets

			if activeFaction != "All" {
				activePlanets = make([]Planet, 0)
				for _, planet := range m.AllPlanets {
					if planet.Initial_owner == activeFaction {
						activePlanets = append(activePlanets, planet)
					}
				}

			}

			m.FilteredPlanets = activePlanets
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
