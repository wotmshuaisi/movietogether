package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	dbPath := "../data/db.sqlite3"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("open db [" + dbPath + "] failed")
		return
	}
	defer func() {
		if err != nil {
			db.Exec("delete from `clients`")
			db.Exec("delete from `message`")
		}
	}()
	// init part
	_, err = db.Exec("CREATE TABLE 'clients'(  'token' Text NOT NULL PRIMARY KEY,  'name' Text NOT NULL, CONSTRAINT 'unique_token' UNIQUE ( 'token' ), CONSTRAINT 'unique_name' UNIQUE ( 'name' ) )")
	if err != nil {
		fmt.Println("init clients error")
		return
	}
	_, err = db.Exec("CREATE TABLE 'messages'(  'datetime' DateTime NOT NULL,  'name' Text NOT NULL,  'message' Text )")
	if err != nil {
		fmt.Println("init messages error")
		return
	}
}
