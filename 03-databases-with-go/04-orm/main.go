package main

import (
	// "fmt"
	"go-mysql/models"
	"go-mysql/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	storage.DB().AutoMigrate(
		&models.Product{},
		&models.InvoiceHeader{},
		&models.InvoiceItem{},
	)

	storage.DB().Model(&models.InvoiceItem{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	storage.DB().Model(&models.InvoiceItem{}).AddForeignKey("invoice_header_id", "invoice_headers(id)", "RESTRICT", "RESTRICT")

	// CREATE PRODUCT
	product1 := models.Product{
		Name: "Go course",
		Price: 74.9,
	}
	obs := "testing with Go"
	product2 := models.Product{
		Name: "Testing course",
		Observations: &obs,
		Price: 149.99,
	}
	product3 := models.Product{
		Name: "Python course",
		Price: 32,
	}

	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
	storage.DB().Create(&product3)

	// // READ ALL DATA
	// products := make([]models.Product, 0)
	// storage.DB().Find(&products)

	// for _, product := range products {
	// 	fmt.Printf("%d - %s\n", product.ID, product.Name)
	// }

	// // READ SINGLE REGISTRY
	// myProduct := models.Product{}
	// storage.DB().First(&myProduct, 2)
	// fmt.Println(myProduct.ID, "-", myProduct.Name)

	// // UPDATE REGISTRY
	// myProduct := models.Product{}
	// myProduct.ID = 1
	// storage.DB().Model(&myProduct).Updates(
	// 	models.Product{Name: "Golang course", Price: 69.99},
	// )

	// // DELETE REGISTRY
	// // Soft Delete
	// myProduct := models.Product{}
	// myProduct.ID = 1
	// storage.DB().Delete(&myProduct)
	// // Hard Delete
	// myProduct := models.Product{}
	// myProduct.ID = 2
	// storage.DB().Unscoped().Delete(&myProduct)

	//////////////////////
	//// TRANSACTIONS ////
	//////////////////////

	// // CREATE
	invoice := models.InvoiceHeader{
		Client: "Albert",
		InvoiceItems: []models.InvoiceItem{
			{ProductID: 1},
			{ProductID: 3},
		},
	}
	storage.DB().Create(&invoice)

}
