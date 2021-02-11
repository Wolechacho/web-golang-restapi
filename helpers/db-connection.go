package helpers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

//CreateDbConnection - connect to mysql database
func CreateDbConnection() {
	fmt.Println("Ready to connect to mysql database .....")

	//the DB returned is actually a pool of underlying DB connections
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/breakingbaddb")
	if err != nil {
		fmt.Printf("Error connecting to DB %+v\n", err)
	} else {
		fmt.Println("Connected to mysql DB")
	}

}

//CreateTestDbConnection - connect to mysql test database
func CreateTestDbConnection() *sql.DB {
	fmt.Println("Ready to connect to mysql database .....")

	//the DB returned is actually a pool of underlying DB connections
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/breakingbaddb")
	if err != nil {
		panic(err)
	} else {
		return db
	}

}
