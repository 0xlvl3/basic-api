package main

import (
	"context"
	"flag"
	"log"

	"github.com/0xlvl3/basic-api/api/db"
	"github.com/0xlvl3/basic-api/api/handles"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// first make a basic api, that has a user, which we can then pull with endpoints.

func main() {

	lp := flag.String("lp", ":8080", "Listening port")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.URI))
	if err != nil {
		log.Fatal(err)
	}

	var (
		// stores
		userStore = db.NewMongoUserStore(client)

		// fiber
		app = fiber.New()
		api = app.Group("/api")

		// handles
		userHandler = handles.NewUserHandler(userStore)
	)

	api.Post("/create", userHandler.HandlePostUser)
	api.Get("/user/:id", userHandler.HandleGetUserByID)

	app.Listen(*lp)
}
