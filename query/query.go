package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type accounting struct {
	GLAccountNumber   int64  `json:"gl_account_number" bson:"gl_account_number"`
	CostCenter        string `json:"cost_center_source" bson:"cost_center_source"`
	DefaultCostCenter int64  `json:"default_cost_center" bson:"default_cost_center"`
}

func main() {
	//client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	client, err := mongo.NewClient("mongodb://admin:securepassword@localhost:29264")
	err = client.Connect(context.TODO())

	// options := options.Find()
	// options.SetLimit(2)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("accounting").Collection("general_ledger")

	cur, err := collection.Find(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cur.Close(context.TODO())

	newFile, err := os.Create("test2.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	for cur.Next(context.TODO()) {
		var elem accounting
		err := cur.Decode(&elem)
		// raw, err := cur.DecodeBytes()
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(elem)
		// fmt.Printf("%v", raw)
		//data := ConvertRawInt64ToString(raw.Lookup("gl_account_number")) + "|" + raw.Lookup("cost_center_source").String() + "|" + ConvertRawInt64ToString(raw.Lookup("default_cost_center")) + "\n"
		data := strconv.FormatInt(elem.GLAccountNumber, 10) + "|" + strconv.FormatInt(elem.DefaultCostCenter, 10) + "|" + elem.CostCenter + "\n"
		_, err = newFile.Write([]byte(data))
		if err != nil {
			log.Fatal(err)
		}
	}

	newFile.Close()

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the connection once no longer needed
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}

}

func ConvertRawInt64ToString(data bson.RawValue) string {
	return strconv.FormatInt(data.Int64(), 10)
}

func ConvertRawToString(data bson.RawValue) string {
	return ""
}
