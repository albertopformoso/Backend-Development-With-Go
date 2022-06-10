package product

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrIDNotFound = errors.New("ID not found")
)

type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}

// Product Model
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %v | %v | %v",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

// Models slice of Model
type Models []*Model

func (m Models) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-20s | %5s | %v | %v\n",
		"id", "name", "observations", "price", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

// Service of product
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used to migrate products
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

// Create is used to create a product
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

// GetAll is used to read all products
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// GetByID is used to read a single product
func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

// Update is used to update a product
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}

	m.UpdatedAt = time.Now()
	return s.storage.Update(m)
}

// Delete is used to delete a product
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
