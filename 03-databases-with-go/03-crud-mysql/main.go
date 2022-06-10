package main

import (
	// "database/sql"
	"fmt"
	"log"

	// "errors"
	// "time"

	// "go-mysql/pkg/invoice"
	// "go-mysql/pkg/invoiceheader"
	// "go-mysql/pkg/invoiceitem"
	"go-mysql/pkg/product"
	"go-mysql/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	storageProduct, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}
	serviceProduct := product.NewService(storageProduct)

	// storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	// storageItems := storage.NewMySQLInvoiceItem(storage.Pool())
	// storageInvoice := storage.NewMySQLInvoice(
	// 	storage.Pool(),
	// 	storageHeader,
	// 	storageItems,
	// )

	// m := &invoice.Model{
	// 	Header: &invoiceheader.Model{
	// 		Client: "Albert",
	// 	},
	// 	Items: invoiceitem.Models{
	// 		&invoiceitem.Model{ProductID: 3},
	// 	},
	// }
	// serviceInvoice := invoice.NewService(storageInvoice)
	// if err := serviceInvoice.Create(m); err != nil {
	// 	log.Fatalf("invoice.Create: %v", err)
	// }

	// CREATE REGISTRY
	m := &product.Model{
		Name:  "Another Course",
		Price: 99.99,
		// Observations: "on fire",
	}

	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}
	fmt.Printf("%+v\n", m)

	// UPDATE REGISTRY
	// m := &product.Model{
	// 	ID: 10,
	// 	Name: "Golang Course",
	// 	Price: 75.50,
	// }
	// err := serviceProduct.Update(m)
	// if err != nil {
	// 	log.Fatalf("product.GetByID: %v", err)
	// }

	// DELETE REGISTRY
	// err := serviceProduct.Delete(2)
	// if err != nil {
	// 	log.Fatalf("product.Delete: %v", err)
	// }

	// GET ALL REGISTRIES
	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	fmt.Println(ms)

	// GET A SINGLE REGISTRY
	// ms, err = serviceProduct.GetByID(3)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	fmt.Println("There's no product with the given id")
	// case err != nil:
	// 	log.Fatalf("product.GetByID: %v", err)
	// default:
	// 	fmt.Println(ms)
	// }

}
