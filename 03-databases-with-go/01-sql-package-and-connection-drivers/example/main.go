package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// MAKE CONNECTION
	db, err := sql.Open("driverName", "dataSourceName") // postgres, connection chain
	if err != nil {
		panic(err)
	}
	defer db.Close() // close the db connection when main finishes its execution

	// Ping indicates that the connection was stablished
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// MANIPULATE DATA
	res, err := db.Exec("INSERT INTO products(name) VALUES ($1)", "Go Course")
	// $1 is a placeholder
	if err != nil {
		log.Fatal(err)
	}

	// LastInsertId is useful to know what ID was assigned to a row that is created
	id, err := res.LastInsertId() // not all DBs support this feature
	if err != nil {
		log.Fatal(err)
	}

	// RowsAffected is useful for counting rows deleted or inserted
	rowsCount, err := res.RowsAffected() // not all DBs support this feature
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("lastInserID: %d, rowsAffected: %d\n", id, rowsCount)


	// OBTAIN DATA COLLECTION
	rows, err := db.Query("SELECT id, name FROM products")
	if err != nil { ... }
	defer rows.Close() // IMPORTANT 

	for rows.Next() {
		// the Go data types must correspond to the data types in the database
		var id uint
		var name string
		if err := rows.Scan(&id, &name); err != nil { // validates the number of arguments, and the supported types
			... 
		}
		fmt.Println(id, name)
	}

	if err := rows.Err(); err != nil { // Ensure that there is no error in the iteration (it is necessary)
		...
	}

	// OBTAIN SINGLE DATA
	product := Product{}
	err := db.QueryRow("SELECT id, name FROM product WHERE id = $1", 6).Scan(
		&product.ID,
		&product.Name,
	)
	switch {
	case err == sql.ErrNoRows:
		log.Print("no hay un producto con este id")
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("Product -> ID: %d, Name: %s", product.ID, product.Name)
	}

	// PREPARED DECLARATIONS
	stmt, err := db.Prepare("INSERT INTO products(name,location) VALUES ($1, $2)")
	if err != nil { ... }
	defer stmt.Close() // IMPORTANTE

	res, err := stmt.Exec("Go Course", "Popayán")
	if err != nil { ... }

	id, err := res.LastInsertId()
	if err != nil { ... }

	rowsAff, err := res.RowsAffected()
	if err != nil { ... }

	res1, err := stmt.Exec("DB course with Go", "Mexico City")
	res2, err := stmt.Exec("Testing course with Go", "Mexico")

	// TRANSACTIONS
	tx, err := db.Begin()
	if err != nil { ... }

	stmtInvoice, err := tx.Prepare("INSERT INTO invoices(client) VALUES(?)")
	if err != nil { tx.Rollback() }
	defer stmtInvoice.Close()

	invRes, err := stmtInvoice.Exec("Albert")
	if err != nil { tx.Rollback() }

	invID, err := invRes.LastInsertId()
	if err != nil { tx.Rollback() }

	stmtItem, err := tx.Prepare("INSERT INTO invoice_items(invoice_id, product, price) VALUES(?,?,?)")
	if err != nil { tx.Rollback() }
	defer stmtItem.Close()

	_, err = stmtItem.Exec(invID, "Go Course", 50)
	if err != nil { tx.Rollback() }

	tx.Commit()

	//NULL VALUES
	type Product struct { Name string}
		for rows.Next() {
		var nameNull sql.NullString
		p := Product{}
		err := rows.Scan(&nameNull)

		p.Name = nameNull.String // code is reduced by omitting validation, since it contains zero value
	}

	type Product struct { Name string }
	for rows.Next() {
		var name *string
		p := Product{}
		err := rows.Scan(&name)
		
		// name contains a value if name is different from nil
		if name != nil { // Here the validation must be done
			p.Name = *name
		}
	}
}

