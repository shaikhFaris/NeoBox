package main

import "github.com/charmbracelet/bubbles/table"

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

	return model{
		choices: []string{"Display Passwords", "Manage Passwords", "Generate Passwords"},
		page:    -1,
		table:   t,
		table2:  t2,
		value:   "",
	}
}
