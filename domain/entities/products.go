package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID
	CategoryId  int
	Name        string
	Description string
	Sku         string
	Price       float64
	Stock       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
}
