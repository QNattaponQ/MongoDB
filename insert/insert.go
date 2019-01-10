package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("test").Collection("trainers")

	// Some dummy data to add to the Database
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	// // Insert a single document
	// insertResult, err := collection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert multiple documents
	trainers := []interface{}{ash, misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	err = client.Disconnect(context.TODO())

	// Add extra options to queries using the options package
	options := options.Find()
	options.SetLimit(2)

	var results []*Trainer

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(context.TODO(), nil, options)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
