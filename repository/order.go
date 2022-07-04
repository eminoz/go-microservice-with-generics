package repository

import (
	"github.com/eminoz/customer-service-with-go/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrderModel model.Order

func (o *Order) InsertOne(ctx *fiber.Ctx, model interface{}) (interface{}, error) {
	insertOne, err := o.Collection.InsertOne(ctx.Context(), model)
	if err != nil {
		return nil, err
	}
	return insertOne.InsertedID, nil

}
func (o *Order) GetOneByID(ctx *fiber.Ctx, id string) (interface{}, error) {
	filter := makeFilter("customerid", id)
	err := o.Collection.FindOne(ctx.Context(), filter).Decode(&OrderModel)

	if err != nil {
		return nil, err
	}
	return OrderModel, nil
}

func (o *Order) UpdateOneById(ctx *fiber.Ctx, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	updateOne, err := o.Collection.UpdateOne(ctx.Context(), filter, update)
	if err != nil {
		return nil, err
	}
	return updateOne, nil
}
