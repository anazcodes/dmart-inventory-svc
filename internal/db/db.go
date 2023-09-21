package db

import (
	"context"
	"log"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(dbUrl string) (*mongo.Database, error) {
	clientOption := options.Client().ApplyURI(dbUrl)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if util.HasError(err) {
		log.Fatalln("failed to connect with db")
	}

	err = client.Ping(context.TODO(), nil)
	if util.HasError(err) {
		log.Fatalln("failed while ping")
	}

	return client.Database("InventorySvc"), nil
}
