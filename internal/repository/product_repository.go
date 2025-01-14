package repository

import "product-api/internal/domain"

type IProductRepository interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id int) error
}
