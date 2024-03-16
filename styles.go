package main

import "github.com/charmbracelet/lipgloss"

var TerminidYellow = lipgloss.Color("#f5cb42")
var AutomatonRed = lipgloss.Color("#b80f00")
var HumanBlue = lipgloss.Color("#008ab8")

var BaseText = lipgloss.NewStyle()

// var BaseText = lipgloss.NewStyle().Width(30).Align(lipgloss.Left)
var TerminidText = BaseText.Copy().Foreground(TerminidYellow)
var AutomatonText = BaseText.Copy().Foreground(AutomatonRed)
var HumanText = BaseText.Copy().Foreground(HumanBlue)

var InfoPanel = lipgloss.NewStyle().Width(50).MarginLeft(50).Background(TerminidYellow)
