package middleware

import (
	"github.com/eminoz/customer-service-with-go/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserValidation() fiber.Handler {
	m := new(model.User)

	validate := validator.New()
	return func(ctx *fiber.Ctx) error {

		ctx.BodyParser(m)
		err := validate.Struct(m)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Next()
	}
}
