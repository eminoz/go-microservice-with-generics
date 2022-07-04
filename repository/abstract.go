package repository

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneEntity[I any](i I) I {
	return i
}

type IBaseEntity interface {
	InsertOne(ctx *fiber.Ctx, model interface{}) (interface{}, error)
	GetOneByID(ctx *fiber.Ctx, id string) (interface{}, error)
	UpdateOneById(ctx *fiber.Ctx, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
}
