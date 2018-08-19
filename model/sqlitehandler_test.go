package model

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Test_dbhandler(t *testing.T) {
	var err error
	dbPath := "../data/db.sqlite3"
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatalf("open db ["+dbPath+"] failed", err)
		return
	}
	handler := NewModel(db)

	defer func() {
		if err != nil {
			db.MustExec("delete from `clients`")
			db.MustExec("delete from `messages`")
		}
		db.Close()
	}()

	err = handler.ClientCreate("1", "1")
	if err != nil {
		t.Fatalf("client create")
	}
	err = handler.ClientCreate("1", "1")
	if err == nil {
		t.Fatalf("client create 1")
	}
	c, err := handler.ClientGet("1")
	if c == nil || err != nil {
		t.Fatalf("client get")
	}
	err = handler.MessageCreate("1", "1")
	if err != nil {
		t.Fatalf("message create")
	}
	result, err := handler.MessagesSelect(1, 15)
	if err != nil || result == nil {
		t.Fatalf("message select")
	}
	return
}
