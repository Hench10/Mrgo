package Mrgo

import (
	"sync"
	"database/sql"
)

type DB struct {
	sync.RWMutex
	db *sql.DB

	search
}
