package repository

import (
	"go-crud/app/config"
	"go-crud/app/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (repository UserRepository) Insert(user domain.User) (domain.User, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	idUser, err := repository.Collection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}

	user.Id = idUser.InsertedID.(primitive.ObjectID)
	return user, nil
}

// func (repository UserRepository) UserExist(username string) bool {
// 	ctx, cancel := config.NewMongoContext()
// 	defer cancel()

// 	var result bson.M
// 	filter := bson.D{{Key: "username", Value: username}}
// 	err := repository.Collection.FindOne(ctx, filter).Decode(&result)
// 	return err == nil
// }

func (repository UserRepository) FindByUsername(username string) (user domain.User, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.D{{Key: "username", Value: username}}
	err = repository.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository UserRepository) FindById(id primitive.ObjectID) (user domain.User, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	err = repository.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository UserRepository) FindAll() (users []domain.User, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	if err != nil {
		return users, err
	}

	for _, document := range documents {
		users = append(users, domain.User{
			Id:       document["_id"].(primitive.ObjectID),
			Name:     document["name"].(string),
			Username: document["username"].(string),
			Password: document["password"].(string),
			Email:    document["email"].(string),
			Role:     document["role"].(string),
		})
	}

	return users, nil
}

func (repository UserRepository) Update(user domain.User) (domain.User, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": user}

	_, err := repository.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository UserRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repository.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
