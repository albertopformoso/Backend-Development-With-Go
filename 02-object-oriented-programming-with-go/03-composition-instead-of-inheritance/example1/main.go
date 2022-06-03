package main

import (
	"fmt"
	"composition/pkg/invoice"
	"composition/pkg/invoiceitem"
	"composition/pkg/customer"
)

func main() {
	i := invoice.New(
		"MX", 
		"CDMX",
		customer.New("Alberto", "St 124 #34", "12345"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Go course", 12.34),
			invoiceitem.New(2, "OOP with Go", 54.23),
			invoiceitem.New(3, "Testing with Go", 90.00),
		},
	)
	i.SetTotal()

	fmt.Printf("%+v\n", i)
}
