package main

import (
	"fmt"
	"log"

	"go-psql/pkg/invoice"
	"go-psql/pkg/invoiceheader"
	"go-psql/pkg/invoiceitem"
	"go-psql/pkg/product"
	"go-psql/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Albert",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 3},
		},
	}
	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

	// CREATE REGISTRY
	// m := &product.Model{
	// 		Name:  "Go Course",
	// 		Price: 70.46,
	// 		// Observations: "on fire",
	// }

	// if err := serviceProduct.Create(m); err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }
	// fmt.Printf("%+v\n", m)

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
