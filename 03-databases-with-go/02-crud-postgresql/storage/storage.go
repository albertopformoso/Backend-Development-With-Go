package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	db      *sql.DB
	once    sync.Once
	connStr string = "postgres://root:secret@localhost:5432/pg_db?sslmode=disable"
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		fmt.Println("Connected to PostgreSQL")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) (null sql.NullString) {
	null = sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return
}

func timeToNull(t time.Time) (null sql.NullTime) {
	null = sql.NullTime{Time: t}
	if null.Time.IsZero() {
		null.Valid = true
	}
	return
}
