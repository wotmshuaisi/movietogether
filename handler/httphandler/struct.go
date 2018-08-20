package httphandler

import (
	"time"
)

type user struct {
	Name string `json:"name"`
}

type result struct {
	Err string `json:"err"`
}

type message struct {
	Name     string    `json:"name"`
	Message  string    `json:"message"`
	Datetime time.Time `json:"datetime"`
}
