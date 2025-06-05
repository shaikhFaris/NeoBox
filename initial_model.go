package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func initialModel() model {
	// models means objects made from structs

	columns := []table.Column{
		{Title: "No", Width: 5},
		{Title: "Service", Width: 12},
		{Title: "Username", Width: 20},
		{Title: "Password", Width: 16}, // Its perfect dont change
	}
	columns2 := []table.Column{
		{Title: "ID", Width: 5},
		{Title: "Service", Width: 10},
	}

	rows := []table.Row{}

	// table styles

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(17),
		table.WithStyles(TableStyles),
	)
	t2 := table.New(
		table.WithColumns(columns2),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
		table.WithStyles(TableStyles2),
	)

	ti := textinput.New()
	ti.Placeholder = "Service"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Cursor.Style = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))
	ti.Cursor.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))
	ti.PlaceholderStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Faint(true)
	ti.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))

	ti2 := textinput.New()
	ti2.Placeholder = "Username"
	ti2.CharLimit = 156
	ti2.Blur()
	ti2.Width = 20
	ti2.Cursor.Style = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))
	ti2.Cursor.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))
	ti2.PlaceholderStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Faint(true)
	ti2.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))

	ti3 := textinput.New()
	ti3.Placeholder = "Password"
	ti3.CharLimit = 156
	ti3.Blur()
	ti3.Width = 20
	ti3.EchoMode = textinput.EchoPassword
	ti3.EchoCharacter = '•'
	ti3.Cursor.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))
	ti3.Cursor.Style = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))
	ti3.PlaceholderStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Faint(true)
	ti3.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))

	ti4 := textinput.New()
	ti4.Placeholder = "ID"
	ti4.CharLimit = 156
	ti4.Blur()
	ti4.Width = 20
	ti4.Cursor.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))
	ti4.Cursor.Style = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))
	ti4.PlaceholderStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Faint(true)
	ti4.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0A0118FF")).Foreground(lipgloss.Color("#0FF74D"))

	return model{
		choices:     []string{"Display Passwords", "Manage Passwords"},
		choices2:    []string{"Create", "Delete"},
		page:        -1,
		table:       t,
		table2:      t2,
		value:       "",
		state:       tableView,
		textInput:   ti,
		textInput2:  ti2,
		textInput3:  ti3,
		textInput4:  ti4,
		err:         nil,
		option:      "",
		createIndex: 0,
		dbOpMsg:     "",
	}
}
