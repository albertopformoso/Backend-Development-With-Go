
# Migrate Products Table

```
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
```

# Migrate Invoice Header Table

```
	storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceheader.Migrate: %v", err)
	}
```

# Migrate Invoice Item Table

```
	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceitem.Migrate: %v", err)
	}
```