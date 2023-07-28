package main

import (
	"fmt"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

)

type stat struct {
	hours int8
}



func main() {
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx,opts)
	if err != nil {
		panic(err)
	}

	db := client.Database("Hours")



}