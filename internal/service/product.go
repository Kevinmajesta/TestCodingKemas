package service

import (
	"errors"
	"kevinmajesta/testkemas/internal/entity"
	"kevinmajesta/testkemas/internal/repository"
)

type ProductService interface {
	CreateProduct(product *entity.Products) (*entity.Products, error)
	UpdateProduct(product *entity.Products) (*entity.Products, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (s *productService) CreateProduct(product *entity.Products) (*entity.Products, error) {
	if product.Name == "" {
		return nil, errors.New("product name cannot be empty")
	}

	if product.Price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}

	if product.Stock <= 0 {
		return nil, errors.New("stock must be greater than 0")
	}

	newProduct := entity.NewProduct(
		product.Name,
		product.Price,
		product.Stock,
	)

	savedProduct, err := s.productRepository.CreateProduct(newProduct)
	if err != nil {
		return nil, err
	}

	return savedProduct, nil
}

func (s *productService) UpdateProduct(product *entity.Products) (*entity.Products, error) {
	if product.Name == "" {
		return nil, errors.New("Product name cannot be empty")
	}

	if product.Price <= 0 {
		return nil, errors.New("Price must be greater than 0")
	}

	if product.Stock <= 0 {
		return nil, errors.New("Stock must be greater than 0")
	}

	updatedProduct, err := s.productRepository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}
