package domain

import "time"

type Product struct {
	ID        int            `json:"Id,omitempty" bson:"_id,omitempty" gorm:"primary_key"`
	Title     string         `json:"Title" bson:"Title" validate:"required"`
	Detail    ProductDetail  `json:"Detail" bson:"Detail" gorm:"-"`
	Prices    []ProductPrice `json:"Prices" bson:"Prices"`
	Inventory int            `json:"Inventory" bson:"Inventory"`
	CreatedAt time.Time      `json:"CreatedAt" bson:"UpdatedAt"`
	UpdatedAt *time.Time     `json:"UpdatedAt" bson:"UpdatedAt" gorm:"index"`
}

// ProductRepository ...
type ProductRepository interface {
	CreateMany(products []Product) (interface{}, error)
	Read() (interface{}, error)
	Update(update Product, id int) (interface{}, error)
	Delete() (interface{}, error)
}

// ProductUsecase ..
type ProductUsecase interface {
	InsertProduct(products []Product) (interface{}, error)
	ListProduct() (interface{}, error)
	UpdateProduct(product Product, id int) (interface{}, error)
	DeleteProduct() (interface{}, error)
}
