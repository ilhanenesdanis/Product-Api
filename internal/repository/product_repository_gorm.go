package repository

import (
	"product-api/internal/domain"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository yeni bir repository oluşturur
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}
