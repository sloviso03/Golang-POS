package services

import (
	"github.com/google/uuid"
	"github.com/sloviso03/Golang-POS/application/dtos"
	"github.com/sloviso03/Golang-POS/application/interfaces"
	"github.com/sloviso03/Golang-POS/domain/entities"
)

type ProductService struct {
	repo interfaces.IProduct
}

func NewProductService(repo interfaces.IProduct) *ProductService {
	return &ProductService{repo: repo}
}

func (productService *ProductService) GetAll() ([]dtos.ProductResponse, error) {
	panic("Not implemented")
}

func (productService *ProductService) GetByID(id uuid.UUID) (*dtos.ProductResponse, error) {
	panic("Not implemented")
}

func (productService *ProductService) Create(request dtos.CreateProductRequest) (*dtos.ProductResponse, error) {
	panic("Not implemented")
}

func (productService *ProductService) Update(id uuid.UUID, request dtos.UpdateProductRequest) (*dtos.ProductResponse, error) {
	panic("Not implemented")
}

func (productService *ProductService) Delete(id uuid.UUID) error {
	return productService.repo.Delete(id)
}

func productMapper(product entities.Product) dtos.ProductResponse {
	return dtos.ProductResponse{
		Id:          product.Id,
		CategoryId:  product.CategoryId,
		Name:        product.Name,
		Description: product.Description,
		Sku:         product.Sku,
		Price:       product.Price,
		Stock:       product.Stock,
		IsActive:    product.IsActive,
	}
}
