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

// imp refrence of mongodb collection
var collection *mongo.Collection

// connect with mongoDB

// / init only runs when applications is started
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

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with ids: ", inserted.InsertedID)
}

// update 1 record

func updateOneMovie(movieId string) {
	id, _ = primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)

}
