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
	items    invoiceitem.Items // 1 to many relationship
}

// New returns a new Invoice
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city: city,
		client: client,
		items: items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}