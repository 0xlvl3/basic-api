package main

import (
	"context"
	"flag"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// first make a basic api, that has a user, which we can then pull with endpoints.

func main() {

	lp := flag.String("lp", ":8080", "Listening port")
	flag.Parse()

	client, err := mongo.Connect(context.TODO, options.Client().ApplyURI(db.URI))
	if err != nil {
		log.Fatal(err)
	}

}
