// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"
)

type User struct {
	ID               int64          `json:"id"`
	Firstname        sql.NullString `json:"firstname"`
	Lastname         sql.NullString `json:"lastname"`
	Username         string         `json:"username"`
	Email            string         `json:"email"`
	Password         string         `json:"password"`
	VerificationCode sql.NullString `json:"verification_code"`
	Verified         sql.NullBool   `json:"verified"`
	CreatedAt        time.Time      `json:"created_at"`
}