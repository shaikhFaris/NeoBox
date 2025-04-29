package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var TableStyles = table.Styles{
	Header: lipgloss.NewStyle().
		Background(lipgloss.Color("#0A0118FF")).
		Foreground(lipgloss.Color("#4949F3")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottomForeground(lipgloss.Color("#4949F3")).
		BorderBottomBackground(lipgloss.Color("#0A0118FF")).
		BorderBottom(true).
		Bold(true).Padding(1, 2),

	Selected: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#4949F3")),

	Cell: lipgloss.NewStyle().Padding(0, 2),
}
var TableStyles2 = table.Styles{
	Header: lipgloss.NewStyle().
		Background(lipgloss.Color("#0A0118FF")).
		Foreground(lipgloss.Color("#4949F3")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottomForeground(lipgloss.Color("#4949F3")).
		BorderBottomBackground(lipgloss.Color("#0A0118FF")).
		BorderBottom(true).
		Bold(true).Padding(0, 2),

	Selected: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#4949F3")),

	Cell: lipgloss.NewStyle().Padding(0, 2),
}
