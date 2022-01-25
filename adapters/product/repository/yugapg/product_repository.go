package product_repository_deprecate

import (
	"gorm.io/gorm"
	"kom-product-service/domain"
)

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) CreateMany(products []domain.Product) (interface{}, error) {
	p.db.
		Model(&domain.Product{}).
		CreateInBatches(&products, 20)

	return nil, p.db.Error
}

func (p *productRepository) Read() (interface{}, error) {
	var products []domain.Product

	p.db.
		Model(&domain.Product{}).
		Find(&products)

	return products, p.db.Error
}

func (p *productRepository) Update(update domain.Product, id int) (interface{}, error) {
	panic("implement me")
}

func (p *productRepository) Delete() (interface{}, error) {
	return nil, nil
}

// New will create an object that represent the domain.BookRepository interface
func New(db *gorm.DB) domain.ProductRepository {
	return &productRepository{db}
}

//// Create ...
//func (mb *mongoBookRepository) CreateMany(databaseName, collectionName string, books []domain.Book) (interface{}, error) {
//
//	newBooks := make([]interface{}, len(books))
//	for i, v := range books {
//		newBooks[i] = v
//	}
//
//	result, err := mb.Conn.Create(databaseName, collectionName, newBooks)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
//
//// Read ...
//func (mb *mongoBookRepository) Read(databaseName, collectionName string, filter interface{}, limit int64, dataModel reflect.Type) (interface{}, error) {
//	return mb.Conn.Read(databaseName, collectionName, filter, limit, dataModel)
//}
//
//// Update ...
//func (mb *mongoBookRepository) Update(databaseName, collectionName string, filter, update interface{}) (interface{}, error) {
//	return mb.Conn.Update(databaseName, collectionName, filter, update)
//}
//
//// Delete ...
//func (mb *mongoBookRepository) Delete(databaseName, collectionName string, filter interface{}) (interface{}, error) {
//	return mb.Conn.Delete(databaseName, collectionName, filter)
//}
