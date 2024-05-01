// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type BookList struct {
	ID     uuid.UUID
	UserID uuid.UUID
	List   json.RawMessage
}

type CompletedTitle struct {
	ID     uuid.UUID
	UserID uuid.UUID
	List   json.RawMessage
}

type GameList struct {
	ID     uuid.UUID
	UserID uuid.UUID
	List   json.RawMessage
}

type MovieList struct {
	ID     uuid.UUID
	UserID uuid.UUID
	List   json.RawMessage
}

type ShowList struct {
	ID     uuid.UUID
	UserID uuid.UUID
	List   json.RawMessage
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	ApiKey    string
}
