package models

import "time"

type Account struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthdat"`
	Sex      string    `json:"sex"`
	Email    string    `json:"email"`
}
