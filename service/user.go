package service

import (
	"github.com/eminoz/customer-service-with-go/model"
	"github.com/eminoz/customer-service-with-go/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserServices interface {
	SaveUser(ctx *fiber.Ctx) (interface{}, error)
	GetUserById(ctx *fiber.Ctx) (interface{}, error)
	UpdateOneUser(ctx *fiber.Ctx) (interface{}, error)
}
type UserService struct {
	UserRepo repository.IBaseEntity
}

func (s *UserService) SaveUser(ctx *fiber.Ctx) (interface{}, error) {
	m := new(model.User)
	ctx.BodyParser(m)
	insertOne, err := s.UserRepo.InsertOne(ctx, m)
	if err != nil {
		return nil, err
	}
	return insertOne, nil
}

func (s *UserService) GetUserById(ctx *fiber.Ctx) (interface{}, error) {
	userId := ctx.Params("id")
	oneByID, err := s.UserRepo.GetOneByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return oneByID, nil
}

func (s *UserService) UpdateOneUser(ctx *fiber.Ctx) (interface{}, error) {
	userID := ctx.Params("id")
	id, _ := primitive.ObjectIDFromHex(userID)
	m := new(model.User)
	ctx.BodyParser(m)
	filter, update := makeFilterAndUpdate("_id", "$set", id, m)

	updateOneById, err := s.UserRepo.UpdateOneById(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateOneById, nil
}
