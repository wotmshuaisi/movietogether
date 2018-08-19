package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	dbPath := "../data/db.sqlite3"
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("open db [" + dbPath + "] failed")
		return
	}
	defer func() {
		if err != nil {
			db.MustExec("delete from `clients`")
			db.MustExec("delete from `message`")
		}
		db.Close()
	}()
	// init part
	_, err = db.Exec("CREATE TABLE 'clients'(  'client_token' Text NOT NULL PRIMARY KEY,  'client_name' Text NOT NULL, CONSTRAINT 'unique_client_token' UNIQUE ( 'client_token' ), CONSTRAINT 'unique_client_name' UNIQUE ( 'client_name' ) )")
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
