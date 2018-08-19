package model

import (
	"github.com/jmoiron/sqlx"
)

// NewModel ...
func NewModel(db *sqlx.DB) MovietogetherDBInterface {
	handler := &dbhandler{
		Db: db,
	}
	return handler
}

// handler struct && interface
type dbhandler struct {
	Db *sqlx.DB
}

// MovietogetherDBInterface ...
type MovietogetherDBInterface interface {
	ClientGet(token string) (*Clients, error)
	ClientCreate(token, name string) error

	MessagesSelect(pagenum, pagesize int) (*MessagesQueryset, error)
	MessageCreate(name, message string) error
}
