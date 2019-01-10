package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")

	// Update a document
	filter := bson.D{{"name", "Ash"}}

	update := bson.D{
		{"$set", bson.D{ //$xxx is operator for MongoDb
			{"age", 10},
			{"city", "Thailand"},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Close the connection once no longer needed
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}

}
