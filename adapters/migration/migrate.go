package migration

import (
	"github.com/arangodb/go-driver"
	"kom-product-service/domain"
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func Migrate(db driver.Database) error {
	var err error

	err = autoMigrate(
		db,
		&domain.Product{},
		&domain.ProductDetail{},
		&domain.ProductPrice{},
	)

	return err
}

func autoMigrate(db driver.Database, models ...interface{}) error {
	for _, model := range models {
		var structName = toSnakeCase(structToString(model))
		if err := createIfNotExist(db, structName); err != nil {
			return err
		}
	}

	return nil
}

func createIfNotExist(db driver.Database, collectionName string) error {
	if ok, _ := db.CollectionExists(nil, collectionName); ok {
		return nil
	}

	_, err := db.CreateCollection(nil, collectionName, &driver.CreateCollectionOptions{})

	return err
}

func structToString(val interface{}) string {
	t := reflect.TypeOf(val)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	return t.Name()
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
