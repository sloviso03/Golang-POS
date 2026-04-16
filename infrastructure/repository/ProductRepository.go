package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sloviso03/Golang-POS/domain/entities"
)

type ProductRepository struct {
	data map[uuid.UUID]entities.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{data: make(map[uuid.UUID]entities.Product)}
}

func (productRepository *ProductRepository) GetAll() ([]entities.Product, error) {
	products := []entities.Product{}

	for _, p := range productRepository.data {
		products = append(products, p)
	}

	return products, nil
}

func (productRepository *ProductRepository) GetByID(id uuid.UUID) (*entities.Product, error) {
	product, ok := productRepository.data[id]

	if !ok {
		return nil, nil
	}
	return &product, nil
}

func (productRepository *ProductRepository) Create(product *entities.Product) error {
	product.Id = uuid.New()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	productRepository.data[product.Id] = *product
	return nil
}

func (productRepository *ProductRepository) Update(product *entities.Product) error {
	_, ok := productRepository.data[product.Id]

	if !ok {
		return errors.New("Product not found")
	}

	product.UpdatedAt = time.Now()
	productRepository.data[product.Id] = *product
	return nil
}

func (productRepository *ProductRepository) Delete(id uuid.UUID) error {
	_, ok := productRepository.data[id]

	if !ok {
		return errors.New("Product not found")
	}

	delete(productRepository.data, id)
	return nil
}
