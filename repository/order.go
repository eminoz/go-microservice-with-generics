package repository

import (
	"github.com/eminoz/customer-service-with-go/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

func (o *Order) UpdateOneById(ctx *fiber.Ctx, id string, update interface{}) (interface{}, error) {
	filter := makeFilter("customerid", id)
	exist := o.IsExist(ctx, filter)
	if exist {
		fil, up := makeFilterAndUpdate("customerid", "$set", id, update)
		updateOne, err := o.Collection.UpdateOne(ctx.Context(), fil, up)
		if err != nil {
			return nil, err
		}
		if updateOne.ModifiedCount == 1 {
			return "order updated", err
		}
		return nil, &fiber.Error{Code: 409, Message: "order already same"}
	}
	return nil, &fiber.Error{Code: 404, Message: "order not found"}

}
func (o *Order) GetAll(ctx *fiber.Ctx) (interface{}, error) {
	find, err := o.Collection.Find(ctx.Context(), bson.D{})
	if err != nil {
		return nil, err
	}
	var result []bson.D
	find.All(ctx.Context(), &result)
	return result, nil
}
func (o Order) IsExist(ctx *fiber.Ctx, filter bson.D) bool {
	documents, _ := o.Collection.CountDocuments(ctx.Context(), filter)
	if documents != 1 {
		return false
	}
	return true
}
