package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Database configuration
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "itscout"
)

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
