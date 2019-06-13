package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHolder interface {
	GetDB() *mongo.Database
}

type Database struct {
	mongoDB *mongo.Database
}

func (db *Database) GetDB() *mongo.Database {
	return db.mongoDB
}

func (db *Database) Close() error {
	return db.mongoDB.Client().Disconnect(nil)
}

func Connect() *Database {
	ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
	fmt.Println("mongourl", os.Getenv("MONGO_URL"))
	clientOpts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		mongoDB: client.Database(os.Getenv("MONGO_DBNAME")),
	}
}
