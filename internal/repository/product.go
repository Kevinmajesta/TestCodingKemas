package repository

import (
	"gorm.io/gorm"
	"kevinmajesta/testkemas/internal/entity"
)

type ProductRepository interface {
	CreateProduct(product *entity.Products) (*entity.Products, error)
	UpdateProduct(product *entity.Products) (*entity.Products, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *entity.Products) (*entity.Products, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) UpdateProduct(product *entity.Products) (*entity.Products, error) {
	fields := make(map[string]interface{})

	if product.Name != "" {
		fields["name"] = product.Name
	}
	if product.Price != 0 {
		fields["price"] = product.Price
	}
	if product.Stock != 0 {
		fields["stock"] = product.Stock
	}

	if err := r.db.Model(product).Where("product_id = ?", product.ProductID).Updates(fields).Error; err != nil {
		return product, err
	}

	return product, nil
}
