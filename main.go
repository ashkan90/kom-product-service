package main

import (
	"context"
	"flag"
	"fmt"
	"kom-product-service/adapters/migration"
	product_http "kom-product-service/adapters/product/delivery/http"
	product_repository "kom-product-service/adapters/product/repository/yugapg"
	product_usecase "kom-product-service/adapters/product/usecase"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	prometheus_http "kom-product-service/adapters/prometheus/handler/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"

	"kom-product-service/config"
)

var (
	cfg    config.Config
	dbConn *gorm.DB
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

	dbUri := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable user=%s password=%s",
		cfg.GetString("database.yugapg.host"),
		cfg.GetInt("database.yugapg.port"),
		cfg.GetString("database.yugapg.name"),
		cfg.GetString("database.yugapg.username"),
		cfg.GetString("database.yugapg.password"),
	)

	var err error
	dbConn, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{
		Logger: logger.New(log.New(), logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      true,
		}),
	})
	if err != nil {
		log.Fatal(err)
	}

	if *migrate {
		go func() {
			log.Info("Migration has been started in another thread...")

			var result = migration.Migrate(dbConn)
			log.Info(fmt.Sprintf("Migration end-up with %s", result))
		}()
	}
}

func main() {
	var e = echo.New()
	e.HideBanner = true

	var productRepository = product_repository.New(dbConn)
	var productUCase = product_usecase.New(productRepository)
	product_http.New(e, productUCase)

	prometheus_http.New(e)

	if cfg.GetBool("debug") {
		e.Start(":" + cfg.GetString("server.port"))
	} else {
		e.Start(":" + os.Getenv("PORT"))
	}
}
