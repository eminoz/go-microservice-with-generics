package repository

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertOneEntity[I IBaseEntity](i I) I {
	return i
}

type IBaseEntity interface {
	InsertOne(ctx *fiber.Ctx, model interface{}) (interface{}, error)
	GetOneByID(ctx *fiber.Ctx, id string) (interface{}, error)
	UpdateOneById(ctx *fiber.Ctx, id string, update interface{}) (interface{}, error)
	GetAll(ctx *fiber.Ctx) (interface{}, error)
	IsExist(ctx *fiber.Ctx, filter bson.D) bool
}
