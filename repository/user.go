package repository

import (
	"github.com/eminoz/customer-service-with-go/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserModel model.User

func (u *User) InsertOne(ctx *fiber.Ctx, model interface{}) (interface{}, error) {

	insertOne, err := u.Collection.InsertOne(ctx.Context(), model)
	if err != nil {
		return nil, err
	}
	return insertOne.InsertedID, nil

}
func (u *User) GetOneByID(ctx *fiber.Ctx, id string) (interface{}, error) {
	userId, _ := primitive.ObjectIDFromHex(id)
	filter := makeFilter("_id", userId)

	err := u.Collection.FindOne(ctx.Context(), filter).Decode(&UserModel)
	if err != nil {
		return nil, err
	}
	return UserModel, nil
}

func (u *User) UpdateOneById(ctx *fiber.Ctx, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	updateOne, err := u.Collection.UpdateOne(ctx.Context(), filter, update)
	if err != nil {
		return nil, err
	}
	return updateOne, nil
}
