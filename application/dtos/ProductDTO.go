package dtos

import "github.com/google/uuid"

type CreateProductRequest struct {
	CategoryId  int     `json:"categoryId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Sku         string  `json:"sku"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
}

type UpdateProductRequest struct {
	CategoryId  int     `json:"categoryId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Sku         string  `json:"sku"`
	Price       float64 `json:"price"`
	Stock       float64 `json:"stock"`
	IsActive    bool    `json:"isActive"`
}

type ProductResponse struct {
	Id          uuid.UUID `json:"id"`
	CategoryId  int       `json:"categoryId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Sku         string    `json:"sku"`
	Price       float64   `json:"price"`
	Stock       float64   `json:"stock"`
	IsActive    bool      `json:"isActive"`
}
