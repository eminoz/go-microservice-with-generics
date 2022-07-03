package repository

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u *GenericCollection) InsertOne(ctx *fiber.Ctx, model interface{}) (interface{}, error) {
	insertOne, err := u.Collection.InsertOne(ctx.Context(), model)
	if err != nil {
		return nil, err
	}
	return insertOne.InsertedID, nil

}
func (u *GenericCollection) GetOneByID(ctx *fiber.Ctx, filter interface{}) (*mongo.SingleResult, error) {
	findOne := u.Collection.FindOne(ctx.Context(), filter)
	return findOne, nil
}