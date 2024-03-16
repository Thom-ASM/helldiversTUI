package main

import (
	"strings"
)

var progressBarChar = "█"

func progressBar(progress float32) string {

	baseString := ""

	initialBarSize := 0
	if progress != 0 {

		initialBarSize = int(float32(width) / (100.0 / progress))
		baseString += HumanText.Render(strings.Repeat(progressBarChar, initialBarSize))
	}
	baseString += TerminidText.Render(strings.Repeat(progressBarChar, width-initialBarSize))

	return baseString + "\n"

}
