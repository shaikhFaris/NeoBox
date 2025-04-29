package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/lib/pq"
	"golang.org/x/term"
)

// let's make a shopping

type model struct {
	choices []string // items on the to-do list
	cursor  int      // which to-do list item our cursor is pointing at
	page    int
	table   table.Model
	value   string // to store the selected row's password for display
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "farisgres@78"
	dbname   = "neobox"
)

type dataFetchedMsg []table.Row

var db *sql.DB

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

func (m model) Init() tea.Cmd {

	// Just return `nil`, which means "no I/O right now, please."
	return nil // in Go nill is same as null in js
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case dataFetchedMsg:
		m.table = table.New(
			table.WithRows(msg),
			table.WithFocused(true),
			table.WithColumns(m.table.Columns()),
			table.WithStyles(TableStyles),
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
				return m, nil
			}
		}
		// show pass view
		if m.page == 0 {
			// fetching and displaying passwords
			// rowsData := fetchData()
			// m.table = table.New(
			// 	table.WithRows(rowsData),
			// 	table.WithFocused(true),
			// 	table.WithColumns(m.table.Columns()),
			// 	table.WithStyles(TableStyles),
			// )

			var cmd tea.Cmd
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "0":
				m.page = -1
			case "2":
				m.page = 1
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
		case "3":
			m.page = 2
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

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
██████╗ ██╗      █████╗  ██████╗██╗  ██╗    ██████╗  ██████╗ ██╗  ██╗
██╔══██╗██║     ██╔══██╗██╔════╝██║ ██╔╝    ██╔══██╗██╔═══██╗╚██╗██╔╝
██████╔╝██║     ███████║██║     █████╔╝     ██████╔╝██║   ██║ ╚███╔╝ 
██╔══██╗██║     ██╔══██║██║     ██╔═██╗     ██╔══██╗██║   ██║ ██╔██╗ 
██████╔╝███████╗██║  ██║╚██████╗██║  ██╗    ██████╔╝╚██████╔╝██╔╝ ██╗
╚═════╝ ╚══════╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚═════╝  ╚═════╝ ╚═╝  ╚═╝
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
		var baseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#0A0118FF")).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#4949F3"))
		var valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#ea00d9"))
		mainContent = baseStyle.Render(m.table.View()) + "\n" + valueStyle.Render(m.value) + "\n\n"
		mainContent += hintStyle.Render("• ↑/↓ navigate  • ↵ reveal password.")
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

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}

func fetchData() []table.Row {
	var rowsData []table.Row
	rows, err := db.Query("SELECT * FROM pass_manager")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var username, service, password string
		err := rows.Scan(&id, &username, &service, &password)
		if err != nil {
			panic(err)
		}
		// m.table=append(m.table.Rows())
		tablerow := table.Row{strconv.Itoa(id), service, username, password}
		rowsData = append(rowsData, tablerow)
	}
	return rowsData
}
func fetchDataCmd() tea.Cmd {
	return func() tea.Msg {
		data := fetchData()
		return dataFetchedMsg(data)
	}
}
