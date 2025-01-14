package usecase

import (
	"product-api/internal/domain"
	"product-api/internal/repository"
)

type ProductUsecase interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id int) (*domain.Product, error)
	CreateProduct(product *domain.Product) error
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id int) error
}

type productUsecase struct {
	repo repository.IProductRepository
}

func (u *productUsecase) GetAllProducts() ([]domain.Product, error) {
	return u.repo.GetAll()
}

func (u *productUsecase) GetProductByID(id int) (*domain.Product, error) {
	return u.repo.GetByID(id)
}
func (u *productUsecase) CreateProduct(product *domain.Product) error {
	return u.repo.Create(product)
}
func (u *productUsecase) UpdateProduct(product *domain.Product) error {
	return u.repo.Update(product)
}
func (u *productUsecase) DeleteProduct(id int) error {
	return u.repo.Delete(id)
}
