package main

import (
	"fmt"
	"math"
	"strings"
)

func planetlist(planets []Planet, height int, paginationIdx int, selectedIndex int) string {

	finalString := ""

	listCount := 0

	for idx, planet := range planets {

		currentPage := int(math.Floor(float64(idx / height)))

		if currentPage > paginationIdx {
			break
		}

		if currentPage == paginationIdx {
			listCount++
			selected := " "
			if idx == selectedIndex {
				selected = ">"
			}

			finalString += applyFactionStyle(fmt.Sprintf("%s %s", selected, planet.Name), planet.Initial_owner) + "\n"

		}

	}

	//pad bottom
	if listCount < height {
		finalString += strings.Repeat("\n", height-listCount)
	}

	finalString += fmt.Sprintf("Page: %d/%d\n", paginationIdx+1, (len(planets)/height)+1)
	return finalString

}
