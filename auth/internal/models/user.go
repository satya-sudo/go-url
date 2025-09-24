package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	Role         string    `db:"role"`
}

// TableName returns the name of the users table.
func (User) TableName() string {
	return "users"
}
