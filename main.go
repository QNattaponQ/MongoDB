package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type restaurant struct {
	ID           string   `json:"_id"`
	Address      address  `json:"address"`
	Borough      string   `json:"borough"`
	Cuisine      string   `json:"cuisine"`
	Grades       []grades `json:"grades"`
	Name         string   `json:"name"`
	RestaurantID string   `json:"restaurant_id"`
}

type address struct {
	Building string    `json:"building"`
	Coord    []float64 `json:"coord"`
	Street   string    `json:"street"`
	Zipcode  string    `json:"zipcode"`
}

type grades struct {
	Date  time.Time `json:"date"`
	Grade string    `json:"grade"`
	Score int       `json:"score"`
}

func main() {

	// Connect DB
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		fmt.Println(err)
	}


	
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	// Query
	resCollection := client.Database("smith").Collection("posnTrn3")

	options := options.Find()
	options.SetLimit(2)
	//options.SetShowRecordID(true)

	//filter := bson.M{"cuisine": "American"}
	cur, err := resCollection.Find(ctx, nil, options)

	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {

		//r := restaurant{}
		//err = cur.Decode(&r)
		raw, err := cur.DecodeBytes()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("===================================================")
		fmt.Printf("%+v\n", raw)
		fmt.Println("===================================================")
	}
}
