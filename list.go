package main

import (
	"fmt"
	"math"
	"strings"
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

			text := fmt.Sprintf("%s %s", selected, planet.Name)

			switch planet.Initial_owner {

			case "Humans":
				finalString += fmt.Sprintf("%s\n", HumanText.Render(text))
			case "Terminids":
				finalString += fmt.Sprintf("%s\n", TerminidText.Render(text))
			case "Automaton":
				finalString += fmt.Sprintf("%s\n", AutomatonText.Render(text))
			}

		}

	}

	//pad bottom
	if len(planets) < height {
		finalString += strings.Repeat("\n", height-len(planets))
	}

	finalString += fmt.Sprintf("Page: %d/%d\n", paginationIdx+1, (len(planets)/height)+1)
	return finalString

}
