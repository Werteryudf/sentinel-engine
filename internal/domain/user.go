package domain

import "time"

type User struct {
	ID           UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsActive     bool
}