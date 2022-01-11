package migration

import (
	"gorm.io/gorm"
	"kom-product-service/domain"
)

func Migrate(g *gorm.DB) error {
	var err error
	err = g.AutoMigrate(
		&domain.Product{},
		&domain.ProductDetail{},
		&domain.ProductPrice{},
	)

	return err
}
