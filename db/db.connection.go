package db

import (
	"context"
	"fmt"
	"log"

	configs "github.com/paulaneesh7/Users_API/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const MONGO_URL = "mongodb+srv://aneesh16117:f5dW0yXL8H9B5is5@cluster0.ors3x.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "UserDB"
const collectionName = "Users"


var Collection *mongo.Collection


// Connect with MongoDB
func init() {
	
	clientOptions := options.Client().ApplyURI(configs.EnvMongoURI())
		
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	Collection = client.Database(dbName).Collection(collectionName)

	
	fmt.Println("Collection instance is Ready✅")
}