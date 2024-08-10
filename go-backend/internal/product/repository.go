package product

import (
	"errors"
)

type Repository interface {
	Save(product *Product) error
	FindByID(id string) (*Product, error)
	FindAll() ([]*Product, error)
	DeleteByID(id string) error
}

type InMemoryRepository struct {
	products map[string]*Product
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		products: make(map[string]*Product),
	}
}

func (r *InMemoryRepository) Save(product *Product) error {
	if product == nil {
		return errors.New("product is nil")
	}
	r.products[product.ID] = product
	return nil
}

func (r *InMemoryRepository) FindByID(id string) (*Product, error) {
	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (r *InMemoryRepository) FindAll() ([]*Product, error) {
	var productList []*Product
	for _, product := range r.products {
		productList = append(productList, product)
	}

	return productList, nil
}

func (r *InMemoryRepository) DeleteByID(id string) error {
	if _, exists := r.products[id]; !exists {
		return errors.New("product not found")
	}

	delete(r.products, id)
	return nil
}
