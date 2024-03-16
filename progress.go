package main

import (
	"strings"
)

var progressBarChar = "█"

func progressBar(progress float32, faction string) string {

	baseString := ""

	initialBarSize := 0
	if progress != 0 {

		initialBarSize = int(float32(width) / (100.0 / progress))
		baseString += HumanText.Render(strings.Repeat(progressBarChar, initialBarSize))
	}
	baseString += applyFactionStyle(strings.Repeat(progressBarChar, width-initialBarSize), faction)

	return baseString + "\n"

}
