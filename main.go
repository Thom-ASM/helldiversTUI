package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var client = &http.Client{Timeout: 10 * time.Second}
var factions [4]string

var height = 20
var width = 20

func (m State) Init() tea.Cmd {
	return fetchAllPlanets_msg
}

func (m State) View() string {
	activeFaction := factions[m.FactionFilterIdx%4]

	s := fmt.Sprintf("Planets (%s)\n", applyFactionStyle(activeFaction, activeFaction))

	s += planetlist(m.FilteredPlanets, 20, m.PaginationIdx, m.SelectedIdx)

	s += progressBar(m.ActivePlanet.Liberation, m.ActivePlanet.Owner)

	s += fmt.Sprintf("%f %% liberated \t %d helldivers liberating planet", m.ActivePlanet.Liberation, m.ActivePlanet.Players)

	return s
}

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case AllPlanetsMsg:
		m.AllPlanets = msg
		m.FilteredPlanets = msg
		m.SelectedIdx = 0

	case SinglePlanetMsg:

		m.ActivePlanet = SinglePlanet(msg)

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			top := m.PaginationIdx * height
			if m.SelectedIdx > top {
				m.SelectedIdx--
			}
		case "down":
			remainder := len(m.FilteredPlanets) - ((m.PaginationIdx) * height)

			bottom := (m.PaginationIdx * height)
			if remainder < height {
				bottom += remainder
			} else {
				bottom += height
			}

			if m.SelectedIdx < bottom-1 {
				m.SelectedIdx++
			}

		case "left":

			maxPages := len(m.FilteredPlanets) / height

			if m.PaginationIdx == 0 {
				m.PaginationIdx = len(m.FilteredPlanets) / height
			} else {
				m.PaginationIdx--
			}

			m.SelectedIdx = m.PaginationIdx % (maxPages + 1) * height

		case "right":
			next := m.PaginationIdx + 1
			maxPages := len(m.FilteredPlanets) / height
			m.PaginationIdx = next % (maxPages + 1)

			m.SelectedIdx = next % (maxPages + 1) * height

		case "enter":

			return m, fetchPlanetInfo(m.FilteredPlanets[m.SelectedIdx].Index)

		case "tab":
			m.FactionFilterIdx++
			m.PaginationIdx = 0

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
