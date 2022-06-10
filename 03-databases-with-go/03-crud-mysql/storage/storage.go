package storage

import (
	"database/sql"
	"fmt"
	"go-mysql/pkg/product"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db      *sql.DB
	once    sync.Once
	mySQLConnStr string = "root:secret@tcp(localhost:3306)/mysql_db?parseTime=true"
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
		db, err = sql.Open("mysql", mySQLConnStr)
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		fmt.Println("Connected to MySQL")
	})
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", postgresConnStr)
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

// DAOProduct factory of product.Storage
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("driver not implemented")
	}
}
