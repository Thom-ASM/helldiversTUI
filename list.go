package main

import (
	"fmt"
	"math"
)

func planetlist(planets []Planet, height int, paginationIdx int, selectedIndex int) string {

	finalString := ""

	for idx, planet := range planets {

		currentPage := int(math.Floor(float64(idx / height)))

		if currentPage > paginationIdx {
			break
		}

		if currentPage == paginationIdx {

			selected := " "
			if idx == selectedIndex {
				selected = ">"
			}

			text := fmt.Sprintf("%s %s\n", selected, planet.Name)

			switch planet.Initial_owner {

			case "Humans":
				finalString += HumanText.Render(text)
			case "Terminids":
				finalString += TerminidText.Render(text)
			case "Automaton":
				finalString += AutomatonText.Render(text)
			}

		}

	}

	finalString += fmt.Sprintf("Page: %d/%d\n", paginationIdx+1, (len(planets)/height)+1)
	return finalString

}
