package api

import (
	"github.com/eminoz/customer-service-with-go/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserServices service.IUserServices
}

func (u *UserController) SaveUser(ctx *fiber.Ctx) error {
	saveUser, err := u.UserServices.SaveUser(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(saveUser)
}
func (u *UserController) GetUserById(ctx *fiber.Ctx) error {
	getUserById, err := u.UserServices.GetUserById(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(getUserById)
}
func (u *UserController) UpdateUser(ctx *fiber.Ctx) error {
	updateOneUser, err := u.UserServices.UpdateOneUser(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(updateOneUser)
}
