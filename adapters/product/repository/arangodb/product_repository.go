package product_repository

import (
	"github.com/arangodb/go-driver"
	"kom-product-service/domain"
)

type productRepository struct {
	col driver.Collection
}

// New will create an object that represent the domain.BookRepository interface
func New(col driver.Collection) domain.ProductRepository {
	return &productRepository{col: col}
}

func (p *productRepository) CreateMany(products []domain.Product) (interface{}, error) {
	var meta, errs, err = p.col.CreateDocuments(nil, products)
	if errs.FirstNonNil() != nil {
		return nil, err
	}

	return meta.Keys(), nil
}

func (p *productRepository) Read() (interface{}, error) {
	var products []domain.Product

	return products, nil
}

func (p *productRepository) Update(update domain.Product, id int) (interface{}, error) {
	panic("implement me")
}

func (p *productRepository) Delete() (interface{}, error) {
	return nil, nil
}
