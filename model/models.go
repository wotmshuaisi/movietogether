package model

import (
	"time"
)

// Clients ...
type Clients struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

// Messages ...
type Messages struct {
	Name     string     `json:"name"`
	Message  string     `json:"message"`
	Datetime *time.Time `json:"datetime"`
}
