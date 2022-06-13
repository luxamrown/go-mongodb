package main

import (
	"context"
	// "encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// ------------------------------------------------------------------------
	// find document

	// coll := client.Database("mohamadelabror-blog").Collection("Blog")
	// id := "1"

	// var result bson.M
	// err = coll.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No post with Id: %s\n", id)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	// ------------------------------------------------------------------------
	// insert document

	// coll := client.Database("mohamadelabror-blog").Collection("Blog")
	// doc := bson.D{{"title", "TESTING | Ini adalah post ketiga"}, {"desc", "Testing, ini adalah deskripsi postingan ketiga, dinject dari golang"}, {"id", "3"}}
	// doc2 := bson.M{"title": "TESTING | Ini adalah post keempat", "desc": "Testing, ini adalah deskripsi postingan keempat, dinject dari golang", "id": "4", "author": "el-abror"}

	// result, err := coll.InsertOne(context.TODO(), doc2)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(result)

	// ------------------------------------------------------------------------
	// update document

	coll := client.Database("mohamadelabror-blog").Collection("Blog")
	id, _ := primitive.ObjectIDFromHex("62a7491c762294322f25bd2e")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"author", "el-abror"}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}
