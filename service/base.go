package service

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertOne(ctx *fiber.Ctx, model interface{}) (interface{}, error)
	GetOneByID(ctx *fiber.Ctx, filter interface{}) (*mongo.SingleResult, error)
}
