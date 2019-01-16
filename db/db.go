package db

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/webcat12345/go-one/conf"
)

func Init() *mongo.Client {
	// connect to the mongo db server
	client, err := mongo.NewClient(conf.MONGO_URL + ":" + conf.MONGO_PORT)
	ctx, _ := context.WithTimeout(context.Background(), conf.MONGO_TIMEOUT)
	err = client.Connect(ctx)

	if err != nil {
		return nil
	} else {
		return client
	}
}
