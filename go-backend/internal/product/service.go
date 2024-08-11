package product

import "errors"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateProduct(name, description, category, location, currency string, price float64) (*Product, error) {
	product := NewProduct(name, description, category, location, currency, price)
	if !product.IsValid() {
		return nil, errors.New("invalid product data")
	}
	err := s.repository.Save(product)
	return product, err
}

func (s *Service) GetProduct(id string) (*Product, error) {
	return s.repository.FindByID(id)
}

func (s *Service) ListProducts() ([]*Product, error) {
	return s.repository.FindAll()
}

func (s *Service) UpdateProduct(id, name, description, category, location, currency string, price float64) (*Product, error) {
	product, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	product.UpdateProduct(name, description, category, location, currency, price)
	err = s.repository.Save(product)
	return product, err
}

func (s *Service) DeleteProduct(id string) error {
	return s.repository.DeleteByID(id)
}
