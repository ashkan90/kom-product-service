package product_usecase

import (
	"kom-product-service/domain"
	"time"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func (p *productUsecase) InsertProduct(products []domain.Product) (interface{}, error) {
	panic("implement me")
}

func (p *productUsecase) ListProduct() (interface{}, error) {
	return []domain.Product{
		{
			ID:        1,
			Title:     "Title field",
			Detail:    domain.ProductDetail{
				Active:    true,
				Barcode:   "TIT-BR-0012",
				Brand:     "Random Brand Name",
				Image:     "the-source-of-image.com",
				Name:      "Brand new product",
			},
			Prices: []domain.ProductPrice{
				{Currency: "TRY", Value: 500.123},
				{Currency: "USD", Value: 65.123},
				{Currency: "EUR", Value: 45.123},
			},
			Inventory: 6,
			CreatedAt: time.Now(),
		},
	}, nil
}

func (p *productUsecase) UpdateProduct(update domain.Product) (interface{}, error) {
	panic("implement me")
}

func (p *productUsecase) DeleteProduct() (interface{}, error) {
	panic("implement me")
}

// New ...
func New(productRepo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		repo: productRepo,
	}
}
//
//func (bu *bookUsecase) InsertBooks(books *[]domain.Book) (interface{}, error) {
//
//	// Dereference the pointer and update the value
//	currentTime := time.Now()
//	for i := range *books {
//		(*books)[i].Created = currentTime
//		(*books)[i].Updated = currentTime
//	}
//
//	result, err := bu.bookRepo.CreateMany(bu.dbName, bu.collName, *books)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
//
//func (bu *bookUsecase) ListBooks(limit int64, dataModel reflect.Type) (interface{}, error) {
//	return bu.bookRepo.Read(bu.dbName, bu.collName, primitive.D{}, limit, dataModel)
//}
//
//func (bu *bookUsecase) UpdateBook(newData domain.Book) (interface{}, error) {
//
//	idPrimitive, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", newData.ID))
//	if err != nil {
//		return nil, err
//	}
//
//	filter := primitive.D{
//		primitive.E{Key: "_id", Value: idPrimitive},
//	}
//
//	product := primitive.D{
//		primitive.E{Key: "$set", Value: primitive.D{
//			primitive.E{Key: "title", Value: newData.Title},
//			primitive.E{Key: "author", Value: newData.Author},
//			primitive.E{Key: "updated", Value: time.Now()},
//		}},
//	}
//
//	return bu.bookRepo.Update(bu.dbName, bu.collName, filter, product)
//}
//
//func (bu *bookUsecase) DeleteBook(bookID string) (interface{}, error) {
//
//	idPrimitive, err := primitive.ObjectIDFromHex(bookID)
//	if err != nil {
//		return nil, err
//	}
//
//	filter := primitive.D{
//		primitive.E{Key: "_id", Value: idPrimitive},
//	}
//
//	result, err := bu.bookRepo.Delete(bu.dbName, bu.collName, filter)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
