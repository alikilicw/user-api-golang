package db

import (
	"database/sql"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db_ *sql.DB) *Store {
	return &Store{
		db:      db_,
		Queries: New(db_),
	}
}
