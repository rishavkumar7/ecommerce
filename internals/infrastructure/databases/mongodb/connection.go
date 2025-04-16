package databases

import (
	"context"

	"github.com/rishavkumar7/ecommerce/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDb(cfg *config.Config) (*mongo.Client, error) {
	mongoUri := cfg.MONGO_URI
	serverApi := options.ServerAPI(options.ServerAPIVersion1)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverApi))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
