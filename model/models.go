package model

import (
	"time"
)

// Clients ...
type Clients struct {
	Token string `db:"client_token"`
	Name  string `db:"client_name"`
}

// Messages ...
type Messages struct {
	Name     string     `db:"name"`
	Message  string     `db:"message"`
	Datetime *time.Time `db:"datetime"`
}

// MessagesQueryset ...
type MessagesQueryset struct {
	TotalPage int         `json:"total_page"`
	TotalNums int         `json:"total_nums"`
	Page      int         `json:"page"`
	Result    []*Messages `json:"result"`
}
