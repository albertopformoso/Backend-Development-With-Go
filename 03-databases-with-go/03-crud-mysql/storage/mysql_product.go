package storage

import (
	"database/sql"
	"fmt"

	"go-mysql/pkg/product"
)

const (
	mySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	mySQLCreateProduct = `INSERT INTO products(name, observations, price, created_at) VALUES(?, ?, ?, ?)`
	mySQLGetAllProduct = `SELECT id, name, observations, price, 
	created_at, updated_at
	FROM products`
	mySQLGetProductByID = mySQLGetAllProduct + " WHERE id = ?"
	mySQLUpdateProduct  = `UPDATE products SET name = ?, observations = ?,
	price = ?, updated_at = ? WHERE id = ?`
	mySQLDeleteProduct = `DELETE FROM products WHERE id = ?`
)

type scanner interface {
	Scan(dest ...interface{}) error
}

// mySQLProduct used for work with mySQL - product
type mySQLProduct struct {
	db *sql.DB
}

// newMySQLProduct return a new pointer of MySQLProduct
func newMySQLProduct(db *sql.DB) *mySQLProduct {
	return &mySQLProduct{db}
}

// Migrate implement the interface product.Storage
func (p *mySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migration completed successfuly!")
	return nil
}

// Create implement the interface product.Storage
func (p *mySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)

	fmt.Printf("Product created successfuly with ID: %d\n", m.ID)
	return nil
}

// GetAll implement the interface product.Storage
func (p *mySQLProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(mySQLGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// GetByID implement the interface product.Storage
func (p *mySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(mySQLGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implement the interface product.Storage
func (p *mySQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("there's no product with ID: %d", m.ID)
	}

	fmt.Println("Product updated successfuly")
	return nil
}

// Delete implement the interface product.Storage
func (p *mySQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(mySQLDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Product deleted successfuly")
	return nil
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
