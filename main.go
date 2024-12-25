package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
// Enter your own MongoDB connection string in the MongoDBLink constant. You can get the connection string from the MongoDB Atlas dashboard.
const MongoDBLink = "mongodb+srv://<username>:<password>@<cluster-url>/<dbname>?retryWrites=true&w=majority"
type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}


var client *mongo.Client


func connectMongoDB() {
	var err error


	connectionString := MongoDBLink
	client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB Atlas!")
}


func handleFormSubmission(c *fiber.Ctx) error {

	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid form data")
	}


	collection := client.Database("testdb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{
		"username": user.Username,
		"password": user.Password,
	})
	if err != nil {
		return c.Status(500).SendString("Failed to save user")
	}

	return c.SendString("User saved successfully!")
}

func main() {

	connectMongoDB()


	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})


	app.Post("/submit", handleFormSubmission)


	log.Fatal(app.Listen(":4000"))
}
