package invoiceitem

// Item contains of information of a invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

// New returns a new Item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value returns the value of the Item Value
func (i *Item) Value() float64 { return i.value }
