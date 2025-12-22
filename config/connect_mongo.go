package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	syncOnce  sync.Once
	db *mongo.Database
)

func ConnectDatabase() *mongo.Database {
	syncOnce.Do(func() {
		fmt.Println("Connecting to mongo cluster")
		ctx := context.Background()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
		if err != nil {
			log.Fatal(err)
		}
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		db = client.Database(os.Getenv("MONGO_DB"))
		fmt.Println("Successfully connected to mongo cluster")
	})

	return db
}
func Database() *mongo.Database {
	if db == nil {
		ConnectDatabase()
	}
	return db
}