package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// for the pass display
	case dataFetchedMsg:
		m.table = table.New(
			table.WithRows(msg),
			table.WithFocused(true),
			table.WithHeight(17), // Its perfect dont change
			table.WithColumns(m.table.Columns()),
			table.WithStyles(TableStyles),
		)
		return m, nil
	case dataFetchedMsg2:
		m.table2 = table.New(
			table.WithRows(msg),
			table.WithFocused(true),
			table.WithHeight(13),
			table.WithColumns(m.table2.Columns()),
			table.WithStyles(TableStyles2),
		)
		return m, nil
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		if m.page == -1 {
			switch msg.String() {
			// The "up" and "k" keys move the cursor up
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			// The "down" and "j" keys move the cursor down
			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

				// The "enter" key and the spacebar (a literal space) toggle
				// the selected state for the item that the cursor is pointing at.
			case "enter":
				m.page = m.cursor
				if m.cursor == 0 {
					return m, fetchDataCmd()
				}
				if m.cursor == 1 {
					return m, fetchDataCmd2()
				}
				return m, nil
			}
		}
		// show pass view
		if m.page == 0 {
			// fetching and displaying passwords
			var cmd tea.Cmd
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "0":
				m.page = -1
			case "2":
				m.page = 1
				return m, fetchDataCmd2()
			case "3":
				m.page = 2
			case "enter":
				row := m.table.SelectedRow()
				if len(row) > 1 {
					m.value = fmt.Sprintf("%s: 👤 %s → 🔐 %s", row[1], row[2], row[3])
				}
				return m, nil

			case "esc":
				if m.table.Focused() {
					m.table.Blur()
				} else {
					m.table.Focus()
				}
			}

			m.table, cmd = m.table.Update(msg)
			return m, cmd
		}

		if m.page == 1 {
			// manager
			var cmd tea.Cmd
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "0":
				m.page = -1
			case "1":
				m.page = 0
				return m, fetchDataCmd()
			case "3":
				m.page = 2

			// case "enter":
			// 	return m, nil
			case "esc":
				if m.table2.Focused() {
					m.table2.Blur()
				} else {
					m.table2.Focus()
				}
			}

			m.table2, cmd = m.table2.Update(msg)
			return m, cmd
		}

		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
		case "0":
			m.page = -1
		case "1":
			m.page = 0
			return m, fetchDataCmd()
		case "2":
			m.page = 1
			return m, fetchDataCmd2()
		case "3":
			m.page = 2
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
