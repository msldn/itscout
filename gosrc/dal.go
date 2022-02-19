package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// Database connections
var DB_USER = os.Getenv("DB_USER")
var DB_PASSWORD = os.Getenv("DB_PASSWORD")
var DB_NAME = os.Getenv("DB_NAME")

// Set up Datbase and open connection
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Execute Query
func exec_query(q string) *sql.Rows {
	db := setupDB()
	rows, err := db.Query(q)
	checkErr(err)
	db.Close()
	return rows
}

// Return all list of CIs in iscout db
func getCis() []Ci {
	// Inint CIs Array
	var CIs []Ci
	rows := exec_query("select * from cis")

	for rows.Next() {
		var id int
		var t string
		var name string
		var created_on string
		err := rows.Scan(&id, &t, &name, &created_on)
		checkErr(err)
		CIs = append(CIs, Ci{ID: id, Type: t, Name: name, Created_on: created_on})
	}
	return CIs
}

// Return all list of CIs in iscout db
func getCi(ids []int) []Ci {
	// Inint CIs Array
	var CIs []Ci
	ids_comma := strings.Trim(strings.Join(strings.Split(fmt.Sprint(ids), " "), ","), "[]")

	rows := exec_query("SELECT * FROM CIS where ID in (" + ids_comma + ")")

	for rows.Next() {
		var id int
		var t string
		var name string
		var created_on string
		err := rows.Scan(&id, &t, &name, &created_on)
		checkErr(err)
		CIs = append(CIs, Ci{ID: id, Type: t, Name: name, Created_on: created_on})
	}
	return CIs
}
