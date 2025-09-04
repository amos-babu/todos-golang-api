package models

import "time"

type Todo struct {
	Id          int       `json:"id" db:"id`
	Name        string    `json:"name" db:"name`
	Description string    `json:"description" db:"description`
	CreatedAt   time.Time `json:"createdAt" db:"createdAt`
}
