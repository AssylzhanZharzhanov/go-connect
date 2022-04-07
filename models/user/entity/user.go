package entity

import "time"

type Users struct {
	ID          int32     `json:"id" db:"id"`
	Photo       string    `json:"photo" db:"photo"`
	Username    string    `json:"username" db:"username"`
	PhoneNumber string    `json:"email" db:"email"`
	Name        string    `json:"name" db:"name"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}