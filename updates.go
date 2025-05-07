package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// you can also get this triggered when the m state is changed
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

	case clearMsg:
		m.dbOpMsg = ""
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
			case "b":
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

			case "ctrl+c":
				return m, tea.Quit

			case "q":
				if m.state != optionsView && m.option != "create" {
					return m, tea.Quit
				}

			case "tab":
				if m.state == tableView {
					m.state = optionsView
					return m, nil
				}
				if m.state == optionsView && m.option == "create" {
					m.createIndex++
					if m.createIndex > 2 {
						m.createIndex = 0
					}

					switch m.createIndex {
					case 0:
						m.textInput.Focus()
						m.textInput2.Blur()
						m.textInput3.Blur()
					case 1:
						m.textInput.Blur()
						m.textInput2.Focus()
						m.textInput3.Blur()
					case 2:
						m.textInput.Blur()
						m.textInput2.Blur()
						m.textInput3.Focus()
					}
					return m, nil
				} else {
					m.state = tableView
					return m, nil
				}
			case "0":
				if m.state != optionsView && m.option != "create" {
					m.page = -1
				}
			case "b":
				if m.state != optionsView && m.option != "create" {
					m.page = -1
				}
			case "1":
				if m.state != optionsView && m.option != "create" {
					m.page = 0
					return m, fetchDataCmd()
				}
			case "3":
				if m.state != optionsView && m.option != "create" {
					m.page = 2
				}

				// for optionsView
			case "up", "k":
				if m.state == optionsView {
					if m.cursor2 > 0 {
						m.cursor2--
					}
				}
			// The "down" and "j" keys move the cursor down
			case "down", "j":
				if m.state == optionsView {
					if m.cursor2 < len(m.choices2)-1 {
						m.cursor2++
					}
				}

			case "enter":
				if m.state == optionsView && m.cursor2 == 0 && m.option == "" {
					m.option = "create"
				}
				if m.state == optionsView && m.cursor2 == 1 && m.option == "" {
					m.option = "delete"
					m.textInput4.Focus()
				}
				if m.state == optionsView && m.option == "create" {
					if len(m.textInput.Value()) != 0 && len(m.textInput2.Value()) != 0 && len(m.textInput3.Value()) != 0 {
						m.formValue[0] = m.textInput.Value()
						m.formValue[1] = m.textInput2.Value()
						m.formValue[2] = m.textInput3.Value()
						m.option = ""
						m.textInput.SetValue("")
						m.textInput2.SetValue("")
						m.textInput3.SetValue("")
						UpdateMsg := insertData(m.formValue[0], m.formValue[1], m.formValue[2])
						m.dbOpMsg = UpdateMsg
						return m, tea.Batch(
							tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
								return clearMsg{}
							}), fetchDataCmd(),
						)
					}
				}

				if m.state == optionsView && m.option == "delete" {
					parsedId, err := strconv.ParseInt(m.textInput4.Value(), 10, 64)
					if len(m.textInput4.Value()) != 0 && err == nil {
						m.option = ""
						m.textInput4.SetValue("")
						deleteMsg := deleteData(int(parsedId))
						m.dbOpMsg = deleteMsg
						return m, tea.Batch(tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
							return clearMsg{}
						}), fetchDataCmd2(),
						)
					}
				}

			// when esc jey is pressed
			case "esc":
				if m.state == tableView {

					if m.table2.Focused() {
						m.table2.Blur()
					} else {
						m.table2.Focus()
					}
				}
				if m.state == optionsView && m.option == "create" {
					m.option = ""
					m.textInput.SetValue("")
					m.textInput2.SetValue("")
					m.textInput3.SetValue("")
					return m, nil
				}
				if m.state == optionsView && m.option == "delete" {
					m.option = ""
					m.textInput4.SetValue("")
					return m, nil
				}
			}

			// when other navigating keys are pressed.
			if m.state == tableView {
				m.table2, cmd = m.table2.Update(msg)
				return m, cmd
			}
			if m.state == optionsView && m.option == "create" {
				switch m.createIndex {
				case 0:
					m.textInput, cmd = m.textInput.Update(msg)
				case 1:
					m.textInput2, cmd = m.textInput2.Update(msg)
				case 2:
					m.textInput3, cmd = m.textInput3.Update(msg)
				}
				return m, cmd
			}
			if m.state == optionsView && m.option == "delete" {
				switch m.createIndex {
				case 0:
					m.textInput4, cmd = m.textInput4.Update(msg)
				}
				return m, cmd
			}
		}

		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit
		case "q":

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
