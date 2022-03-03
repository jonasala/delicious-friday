package db

import (
	"context"
	"time"

	mongo_driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo_driver.Client
var DB *mongo_driver.Database

func Connect(mongoURI, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo_driver.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	if err := Client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	DB = Client.Database(dbName)
	return nil
}
