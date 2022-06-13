package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://<username>:<password>@<host>:<port>/?retryWrites=true&w=majority"

// if your password contain '@', change '@' with '%40'

func main() {
	// Create new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Ping the primary

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	//succes to connect

	fmt.Println("Connect Success")

}
