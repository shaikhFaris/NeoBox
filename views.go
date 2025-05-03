package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func (m model) View() string {

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return "Error getting terminal size."
	}

	hintStyle := lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Height(1).Faint(true).Background(lipgloss.Color("#0A0118FF"))
	hooklinestyle := lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Height(1).
		Foreground(lipgloss.Color("#0FF74D")).Background(lipgloss.Color("#0A0118FF"))
	mainContent := ""

	if m.page == -1 {

		mainContent +=
			`
███╗   ██╗███████╗ ██████╗     ██████╗  ██████╗ ██╗  ██╗
████╗  ██║██╔════╝██╔═══██╗    ██╔══██╗██╔═══██╗╚██╗██╔╝
██╔██╗ ██║█████╗  ██║   ██║    ██████╔╝██║   ██║ ╚███╔╝ 
██║╚██╗██║██╔══╝  ██║   ██║    ██╔══██╗██║   ██║ ██╔██╗ 
██║ ╚████║███████╗╚██████╔╝    ██████╔╝╚██████╔╝██╔╝ ██╗
╚═╝  ╚═══╝╚══════╝ ╚═════╝     ╚═════╝  ╚═════╝ ╚═╝  ╚═╝
`
		mainContent += hooklinestyle.Render("\nA quiet vault where your passwords rest, waiting for the right command.")
		mainContent += "\n\n"
		maxLen := 0
		for _, choice := range m.choices {
			if len(choice) > maxLen {
				maxLen = len(choice)
			}

		}

		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			// %-*s ensures padding for left alignment
			mainContent += fmt.Sprintf("%s %d. %-*s\n", cursor, i+1, maxLen, choice)
		}

		mainContent += hintStyle.Render("\n• ↑/↓ navigate  • ↵ select  • q quit\n")
		// mainContent += ""

	}

	if m.page == 0 {
		var baseTableStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#0A0118FF")).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#4949F3")).MarginTop(2)

		var valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#ea00d9"))

		mainContent = baseTableStyle.Render(m.table.View()) + "\n" + valueStyle.Render(m.value) + "\n\n"
		mainContent += hintStyle.Render("• ↑/↓ navigate  • ↵ reveal password  • b back")
	}

	if m.page == 1 {
		// var baseTableStyle = lipgloss.NewStyle().
		// 	Foreground(lipgloss.Color("#ffffff")).
		// 	Background(lipgloss.Color("#0A0118FF"))
		maxLen := 0
		options := ""
		for _, choice := range m.choices2 {
			if len(choice) > maxLen {
				maxLen = len(choice)
			}
		}

		for i, choice := range m.choices2 {
			cursor := " "
			if m.cursor2 == i {
				cursor = ">"
			}
			// cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#0FF74D")).Bold(true).Background(lipgloss.Color("#0A0118FF"))
			// OptionTextStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Bold(true).Background(lipgloss.Color("#0A0118FF"))

			// %-*s ensures padding for left alignment
			options += fmt.Sprintf("%s %d. %-*s\n\n", cursor, i+1, maxLen, choice)
		}
		optionsStyle := lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Height(14).AlignVertical(lipgloss.Center).PaddingLeft(6).PaddingRight(7).Foreground(lipgloss.Color("#ffffff"))
		options = optionsStyle.Render(options)

		if m.option == "create" {
			// m.textInput
			options = fmt.Sprintf(
				"Enter your Details\n\n%s\n\n%s\n\n%s\n\n%s",
				m.textInput.View(),
				m.textInput2.View(),
				m.textInput3.View(),
				lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Faint(true).Render("• tab: switch • esc: exit\n• enter: submit"),
			)
			options = lipgloss.NewStyle().PaddingLeft(10).Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#ffffff")).Height(14).Render(options)
		}

		var baseTableStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#0A0118FF"))
		if m.state == tableView {
			mainContent = lipgloss.JoinHorizontal(lipgloss.Top, baseTableStyle.Render(m.table2.View()), options)
		} else {
			mainContent = lipgloss.JoinHorizontal(lipgloss.Top, baseTableStyle.Render(m.table2.View()), options)
		}
		mainContent = lipgloss.NewStyle().Padding(2).PaddingBottom(1).MarginTop(4).PaddingRight(3).Background(lipgloss.Color("#0A0118FF")).BorderStyle(lipgloss.RoundedBorder()).BorderBackground(lipgloss.Color("#0A0118FF")).BorderForeground(lipgloss.Color("#0FF74D")).Render(mainContent)
		mainContent += "\n"

		mainContent += hintStyle.Render("\n• tab: switch focus • b: back • q: exit\n")

	}

	// Centered main content (height - 1 to leave room for footer)
	centeredStyle := lipgloss.NewStyle().
		Width(width).
		Height(height - 1). // Reserve space for footer
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Foreground(lipgloss.Color("#ffffffff")).
		Background(lipgloss.Color("#0A0118FF"))
	mainContent = centeredStyle.Render(mainContent)

	// Footer (aligned to bottom-left)
	pageCounter := strconv.Itoa(m.page + 1)

	pageCounterStyle := lipgloss.NewStyle().
		Width(width).
		Height(1).
		AlignHorizontal(lipgloss.Left).
		Padding(0, 1).
		Background(lipgloss.Color("#0A0118FF"))
	pageCounter = pageCounterStyle.Render(pageCounter)

	return mainContent + "\n" + pageCounter

}
