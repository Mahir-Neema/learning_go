package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://mahirneema6:N5xKAjBsmXAPSJeq@cluster1.evpa69u.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

// imp
var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connection to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb Connection Success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("collection instance is ready")
}
