package domain

import "time"

type ProductDetail struct {
	ID        int        `json:"-,omitempty" bson:"_id,omitempty" gorm:"primary_key"`
	ProductID int        `json:"-" bson:"ProductId" gorm:"foreign_key"`
	Active    bool       `json:"Active" bson:"active"`
	Barcode   string     `json:"Barcode" bson:"barcode"`
	Brand     string     `json:"Brand" bson:"brand"`
	Image     string     `json:"Image" bson:"image"`
	Name      string     `json:"Name" bson:"name"`
	UpdatedAt *time.Time `json:"UpdatedAt" bson:"UpdatedAt" gorm:"index"`
}
