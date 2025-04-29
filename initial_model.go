package main

import "github.com/charmbracelet/bubbles/table"

func initialModel() model {
	// models means objects made from structs

	columns := []table.Column{
		{Title: "No", Width: 5},
		{Title: "Service", Width: 12},
		{Title: "Username", Width: 20},
		{Title: "Password", Width: 20},
	}

	rows := []table.Row{}

	// table styles

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithStyles(TableStyles),
	)

	return model{
		choices: []string{"Display Passwords", "Manage Passwords", "Generate Passwords"},
		page:    -1,
		table:   t,
		value:   "",
	}
}
