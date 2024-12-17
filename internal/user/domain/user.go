package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Document     string             `json:"document" bson:"document"`
	DocumentType string             `json:"documentType" bson:"documentType"`
	Phone        string             `json:"phone" bson:"phone"`
	Password     string             `json:"password" bson:"password"`
}
