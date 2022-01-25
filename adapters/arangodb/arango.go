package arangodb

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	log "github.com/sirupsen/logrus"
	"os"
)

type ArangoOptions struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

func ArangoUtilize(opt ArangoOptions) (driver.Database, error) {
	var err error
	var dbConn driver.Database

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{fmt.Sprintf("http://%s:%d", opt.Host, opt.Port)},
	})
	if err != nil {
		log.Fatal("something went wrong while opening connection...")
		return nil, err
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(opt.Username, opt.Password),
	})
	if err != nil {
		log.Fatal("something went wrong while creating client...")
		return nil, err
	}

	if ok, _ := client.DatabaseExists(nil, opt.Database); !ok {
		// Open {database} database
		dbConn, err = client.CreateDatabase(nil, opt.Database, &driver.CreateDatabaseOptions{})
		if err != nil {
			return nil, err
		}
	} else {
		dbConn, err = client.Database(nil, opt.Database)
	}

	return dbConn, nil
}

func ArangoOpenCollection(db driver.Database, ctx context.Context, colName string) driver.Collection {
	var col, err = db.Collection(ctx, colName)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	return col
}
