package helpers

import (
	"fmt"
	"log"
	"time"

	"github.com/paulaneesh7/Users_API/db"
	"github.com/paulaneesh7/Users_API/model"
	"github.com/paulaneesh7/Users_API/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user model.User) {

	ctx, cancel := utils.CreateContext(10 * time.Second)
	defer cancel()

	res, err := db.Collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("User created successfully✅: ", res)
}

func GetAllUsers() []model.User {

	ctx, cancel := utils.CreateContext(10 * time.Second)
	defer cancel()

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



func GetUserById(id string) model.User {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := utils.CreateContext(10 * time.Second)
	defer cancel()

	filter := bson.M{"_id": Id}

	var user model.User
	err = db.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("User found✅: ", user)
	return user
}


// UpdateById can also be istead of UpdateOne (check the mongo docs for UpdateById)
func UpdateUserById(id string, user model.User) model.User {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := utils.CreateContext(10 * time.Second)
	defer cancel()

	filter := bson.M{"_id": Id}
	update := bson.M{"$set": bson.D{{"name", user.Name}, {"email", user.Email}, {"gender", user.Gender}, {"education", user.Education}, {"bio", user.Bio}}}

	_, err = db.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Retrieve the updated user
	var updatedUser model.User
	err = db.Collection.FindOne(ctx, filter).Decode(&updatedUser)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("User updated successfully✅: ", updatedUser)
	return updatedUser

}

func DeleteUserById(id string) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := utils.CreateContext(10 * time.Second)
	defer cancel()

	filter := bson.M{"_id": Id}
	_, err = db.Collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("User deleted successfully✅ ")
}



func GetCountOfUsers() int64 {
	ctx, cancel := utils.CreateContext(10 * time.Second)
	defer cancel()

	count, err := db.Collection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err.Error())
	}

	return count
}