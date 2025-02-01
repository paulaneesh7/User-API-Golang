package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/paulaneesh7/Users_API/db"
	"github.com/paulaneesh7/Users_API/model"
	"go.mongodb.org/mongo-driver/bson"
)


func CreateUser(user model.User) {

	res, err := db.Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("User created successfully✅: ", res)
}


func GetAllUsers() []model.User {

	ctx := context.Background()

	cursor, err := db.Collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err.Error())
	}

	var results []model.User

	for cursor.Next(ctx) {
		var result model.User
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err.Error())
		}
		results = append(results, result)
	}

	defer cursor.Close(ctx)
	fmt.Println("All users: ", results)

	return results
}
