package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db              *gorm.DB
	once            sync.Once
	mySQLConnStr    string = "root:secret@tcp(localhost:3306)/mysql_db?parseTime=true"
	postgresConnStr string = "postgres://root:secret@localhost:5432/pg_db?sslmode=disable"
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// creates the connection with the db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("mysql", mySQLConnStr)
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		fmt.Println("Connected to MySQL")
	})
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("postgres", postgresConnStr)
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		fmt.Println("Connected to PostgreSQL")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB {
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
