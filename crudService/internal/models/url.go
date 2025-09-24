package models

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ShortCode    string     `db:"short_code" json:"shortCode"`
	LongURL      string     `db:"long_url" json:"longUrl"`
	UserID       uuid.UUID  `db:"user_id" json:"userId"`
	CreatedAt    time.Time  `db:"created_at" json:"createdAt"`
	ExpirationAt *time.Time `db:"expiration_at,omitempty" json:"expirationAt,omitempty"`
	Hits         int64      `db:"hits" json:"hits"`
}
