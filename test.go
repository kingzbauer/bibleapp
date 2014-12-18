package main

import (
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"fmt"
	"github.com/droundy/goopt"
)

var dbConn *sqlite3.Conn

// Variables passed in at the terminal to be extracted using goopt
// var info : the table to view information about
var info = goopt.Alternatives([]string{"-i", "--info"},
	[]string{"verses", "chapters", "metadata", "annotations"},
	"View information about the database table ")

var dbName = goopt.String([]string{"--db", "--database"}, "amp.sqlite3",
	"The name of the database file")


/*
 * dbInit initializes or opens a database connection storing in the @dbConn variable
 */

func DbInit(){
	// Initialize the database connection
	_dbConn, err := sqlite3.Open(*dbName)
	if err != nil{
		fmt.Print("\033[31m")
		fmt.Println("Error in connecting to the db")
		panic(err)
	}
	dbConn = _dbConn
}

func ain(){
	goopt.Description = func() string {
		return "A bible app test"
	}
	
	goopt.Version = "0.1"
	goopt.Summary = "Bible application by golang"
	goopt.Parse(nil)

	DbInit()
	results := query("PRAGMA table_info(chapters);")
	fmt.Print("\033[34m")
	for _, v := range results{
		for key, value := range v{
			fmt.Print(fmt.Sprintf(
				"%10s :: %s", key, value))
		}
		fmt.Println("")
		fmt.Println("=======================================================================================================")
	}
	fmt.Print("\033[0m")
	
}

func query(sqlQuery string) (results []sqlite3.RowMap) {
	// Execute query
	for s, err := dbConn.Query(sqlQuery); err == nil; err = s.Next(){
		row := make(sqlite3.RowMap)
		// Scan the current row
		s.Scan(row)
		// Append to results
		results = append(results, row)
	}
	return
}
