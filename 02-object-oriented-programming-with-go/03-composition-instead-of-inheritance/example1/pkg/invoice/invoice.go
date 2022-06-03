package invoice

import (
	"composition/pkg/customer"
	"composition/pkg/invoiceitem"
)

// Invoice is the struct of an invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer  // 1 to 1 relationship
	items    []invoiceitem.Item // 1 to many relationship
}

// New returns a new Invoice
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city: city,
		client: client,
		items: items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}