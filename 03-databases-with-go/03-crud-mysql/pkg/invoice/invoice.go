package invoice

import (
	"go-mysql/pkg/invoiceheader"
	"go-mysql/pkg/invoiceitem"
)

// MOdel of invoice
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Storage interface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of invoice
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create a new invoice
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
