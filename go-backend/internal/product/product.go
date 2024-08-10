package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Currency    string    `json:"currency"`
	Category    string    `json:"category"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func generateID() string {
	return uuid.New().String()
}

func NewProduct(name, description, category, location, currency string, price float64) *Product {
	return &Product{
		ID:          generateID(),
		Name:        name,
		Description: description,
		Price:       price,
		Currency:    currency,
		Category:    category,
		Location:    location,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (p *Product) UpdateProduct(name, description, category, location, currency string, price float64) {
	p.Name = name
	p.Description = description
	p.Category = category
	p.Location = location
	p.Currency = currency
	p.Price = price
}

func (p *Product) IsValid() bool {
	return p.Name != "" && p.Price >= 0
}
