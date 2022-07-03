package api

import "github.com/gofiber/fiber/v2"

type UserService interface {
	SaveUser(ctx *fiber.Ctx) (interface{}, error)
	GetUserById(ctx *fiber.Ctx) (interface{}, error)
}
type UserController struct {
	UserServices UserService
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
