package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/arangodb/go-driver"
	"kom-product-service/adapters/arangodb"
	"kom-product-service/adapters/migration"
	product_http "kom-product-service/adapters/product/delivery/http"
	product_repository "kom-product-service/adapters/product/repository/arangodb"
	product_usecase "kom-product-service/adapters/product/usecase"

	prometheus_http "kom-product-service/adapters/prometheus/handler/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"

	"kom-product-service/config"
)

var (
	cfg    config.Config
	dbConn driver.Database
	ctx    = context.Background()
)

func init() {
	var migrate = flag.Bool("migrate", false, "Would you like to migrate the database ?")

	flag.Parse()

	log.Info(fmt.Sprintf("Migration mode: %t", *migrate))

	cfg = config.New()
	if cfg.GetBool("debug") {
		log.Info("Debug mode: true")
	}

	var err error
	dbConn, err = arangodb.ArangoUtilize(arangodb.ArangoOptions{
		Host:     cfg.GetString("database.arangodb.host"),
		Port:     cfg.GetInt("database.arangodb.port"),
		Database: cfg.GetString("database.arangodb.name"),
		Username: cfg.GetString("database.arangodb.username"),
		Password: cfg.GetString("database.arangodb.password"),
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	if *migrate {
		log.Info("Migration has been started...")

		var result = migration.Migrate(dbConn)
		log.Info(fmt.Sprintf("Migration end-up with %s", result))
	}
}

func main() {
	var e = echo.New()
	e.HideBanner = true

	var productRepository = product_repository.New(arangodb.ArangoOpenCollection(dbConn, nil, "product"))
	var productUCase = product_usecase.New(productRepository)
	product_http.New(e, productUCase)

	prometheus_http.New(e)

	if cfg.GetBool("debug") {
		e.Start(":" + cfg.GetString("server.port"))
	} else {
		e.Start(":" + os.Getenv("PORT"))
	}
}
