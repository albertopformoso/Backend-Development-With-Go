package storage

import (
	"database/sql"
	"fmt"
	"go-psql/pkg/invoiceheader"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS "invoice_headers" (
        "id" SERIAL NOT NULL,
        "client" VARCHAR NOT NULL,
        "created_at" TIMESTAMP NOT NULL DEFAULT now(),
        "updated_at" TIMESTAMP,
        CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id));`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client)
		VALUES ($1) RETURNING id, created_at;`
)

// psqlInvoiceHeader used to work with postgres - invoice header
type psqlInvoiceHeader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader return a new pointer of psqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *psqlInvoiceHeader {
	return &psqlInvoiceHeader{db}
}

// Migrate implements the interface invoiceheader.Storage
func (p *psqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Invoice Header migration executed sucessfuly!")
	return nil
}

// CreateTX implement the interface invoiceheader.Storage
func (p *psqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
