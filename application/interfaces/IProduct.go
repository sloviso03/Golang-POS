package interfaces

import (
	"github.com/google/uuid"
	"github.com/sloviso03/Golang-POS/domain/entities"
)

type Product = entities.Product

type IProduct interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (*Product, error)
	Create(p *Product) error
	Update(p *Product) error
	Delete(id uuid.UUID) error
}
