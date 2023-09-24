package db

import (
	"context"
	"log"
	"time"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(dbUrl string) (*mongo.Database, error) {
	clientOption := options.Client().ApplyURI(dbUrl)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOption)
	if util.HasError(err) {
		log.Fatalln("failed to connect with db")
	}

	err = client.Ping(context.TODO(), nil)
	if util.HasError(err) {
		log.Fatalln("failed while ping")
	}

	return client.Database("InventorySvc"), nil
}
