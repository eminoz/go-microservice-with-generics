package service

import (
	"github.com/eminoz/customer-service-with-go/model"
	"github.com/eminoz/customer-service-with-go/repository"
	"github.com/gofiber/fiber/v2"
)

type IOrderService interface {
	SaveOrder(ctx *fiber.Ctx) (interface{}, error)
	GetOrderById(ctx *fiber.Ctx) (interface{}, error)
	UpdateOneOrder(ctx *fiber.Ctx) (interface{}, error)
}
type OrderService struct {
	OrderRepo repository.IBaseEntity
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
	getOneByID, err := o.OrderRepo.GetOneByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return getOneByID, nil
}
func (o *OrderService) UpdateOneOrder(ctx *fiber.Ctx) (interface{}, error) {
	customerid := ctx.Params("id")
	order := new(model.Order)
	ctx.BodyParser(order)
	filter, update := makeFilterAndUpdate("customerid", "$set", customerid, o)
	updateOneById, err := o.OrderRepo.UpdateOneById(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateOneById.ModifiedCount, nil
}
