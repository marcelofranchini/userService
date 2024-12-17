package infra

import (
	"context"
	"userService/internal/user/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCreateMongoRepository struct {
	collection *mongo.Collection
}

func UserCreateRepository(db *mongo.Database) *UserCreateMongoRepository {
	return &UserCreateMongoRepository{
		collection: db.Collection("users"),
	}
}

func (repo *UserCreateMongoRepository) Save(user *domain.User) error {
	_, err := repo.collection.InsertOne(context.Background(), user)
	return err
}

func (repo *UserCreateMongoRepository) FindByID(id string) (*domain.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user domain.User
	err = repo.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
