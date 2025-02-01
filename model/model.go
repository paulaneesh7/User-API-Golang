package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Gender    string             `json:"gender"`
	Education string             `json:"education"`
	Bio       string             `json:"bio"`
}
