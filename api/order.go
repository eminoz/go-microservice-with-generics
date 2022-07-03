package api

import "github.com/gofiber/fiber/v2"

type OrderService interface {
	SaveOrder(ctx *fiber.Ctx) (interface{}, error)
	GetOrderById(ctx *fiber.Ctx) (interface{}, error)
}
type OrderController struct {
	OrderServices OrderService
}

func (o *OrderController) SaveOrder(ctx *fiber.Ctx) error {
	saveOrder, err := o.OrderServices.SaveOrder(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(saveOrder)
}

func (o *OrderController) GetOrderById(ctx *fiber.Ctx) error {
	getOrderById, err := o.OrderServices.GetOrderById(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(getOrderById)
}
