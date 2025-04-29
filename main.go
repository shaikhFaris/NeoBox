package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"
)

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

func (m model) Init() tea.Cmd {

	// Just return `nil`, which means "no I/O right now, please."
	return nil // in Go nill is same as null in js
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
