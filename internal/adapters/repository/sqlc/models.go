// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repository

import (
	"time"
)

type User struct {
	ID             int64
	CreatedAt      time.Time
	Email          string
	Phone          string
	HashedPassword string
}
