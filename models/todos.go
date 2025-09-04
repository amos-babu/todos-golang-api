package models

import "time"

type Todo struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Desc string    `json:"desc"`
	Time time.Time `json:"time"`
}
