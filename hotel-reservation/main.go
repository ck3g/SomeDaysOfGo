package main

import (
	"context"
	"flag"
	"log"

	"github.com/ck3g/SomeDaysOfGo/hotel-reservation/api"
	"github.com/ck3g/SomeDaysOfGo/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi          = "mongodb://localhost:27017"
	dbname         = "hotel-reservation"
	userCollection = "users"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	// ctx := context.Background()
	// collection := client.Database(dbname).Collection(userCollection)
	// user := types.User{
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// }
	// res, err := collection.InsertOne(ctx, user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res.InsertedID)

	listenAddr := flag.String("listenAddr", ":5005", "Server listen address of API")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/users/:id", userHandler.HandleGetUser)
	apiv1.Post("/users", userHandler.HandlePostUser)
	app.Listen(*listenAddr)
}
