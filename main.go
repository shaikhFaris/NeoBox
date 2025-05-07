package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"
)

type model struct {
	choices     []string // items on the to-do list
	cursor      int      // which to-do list item our cursor is pointing at
	page        int
	table       table.Model
	table2      table.Model
	value       string // to store the selected row's password for display
	state       sessionState
	cursor2     int
	choices2    []string
	option      string
	textInput   textinput.Model
	textInput2  textinput.Model
	textInput3  textinput.Model
	textInput4  textinput.Model
	err         error
	createIndex int
	formValue   [3]string
	dbOpMsg     string
}

const (
	host                   = "localhost"
	port                   = 5432
	user                   = "postgres"
	password               = "farisgres@78"
	dbname                 = "neobox"
	tableView sessionState = iota
	optionsView
)

type dataFetchedMsg []table.Row
type dataFetchedMsg2 []table.Row
type sessionState uint
type clearMsg struct{}

// type errMsg error

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
func fetchData2() []table.Row {
	var rowsData []table.Row
	rows, err := db.Query("SELECT id, service FROM pass_manager")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var service string
		err := rows.Scan(&id, &service)
		if err != nil {
			panic(err)
		}
		// m.table=append(m.table.Rows())
		tablerow := table.Row{strconv.Itoa(id), service}
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
func fetchDataCmd2() tea.Cmd {
	return func() tea.Msg {
		data := fetchData2()
		return dataFetchedMsg2(data)
	}
}

func insertData(s1 string, s2 string, s3 string) string {
	query := fmt.Sprintf(`INSERT INTO pass_manager (username, service, password) VALUES ('%s', '%s', '%s');`, s1, s2, s3)
	_, err := db.Query(query)
	if err != nil {
		// panic(err)
		return "failed to update DB"
	}
	return "Added\nsuccessfully"
}

func deleteData(id int) string {
	query := fmt.Sprintf(`DELETE FROM pass_manager WHERE id=%d`, id)
	_, err := db.Query(query)
	if err != nil {
		return "error\n(might be because this id does'nt exist)"
	}
	return "deleted\nsuccessfully"
}
