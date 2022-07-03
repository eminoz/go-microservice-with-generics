package service

import (
	"github.com/eminoz/customer-service-with-go/model"
	"github.com/gofiber/fiber/v2"
)

type OrderService struct {
	OrderRepo Repository
}

func (o *OrderService) SaveOrder(ctx *fiber.Ctx) (interface{}, error) {
	m := new(model.Order)
	ctx.BodyParser(m)
	insertOne, err := o.OrderRepo.InsertOne(ctx, m)
	if err != nil {
		return nil, err
	}
	return insertOne, nil
}

func (o *OrderService) GetOrderById(ctx *fiber.Ctx) (interface{}, error) {
	userId := ctx.Params("id")
	var model model.Order
	filter := makeFilter("customerid", userId)
	getOneByID, err := o.OrderRepo.GetOneByID(ctx, filter)
	getOneByID.Decode(&model)
	if err != nil {
		return nil, err
	}
	return model, nil
}
