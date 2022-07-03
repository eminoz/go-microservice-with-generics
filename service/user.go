package service

import (
	"github.com/eminoz/customer-service-with-go/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	UserRepo Repository
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
	id, err2 := primitive.ObjectIDFromHex(userId)
	if err2 != nil {
		return nil, err2
	}

	filter := makeFilter("_id", id)
	var model model.User
	getOneByID, err := s.UserRepo.GetOneByID(ctx, filter)
	getOneByID.Decode(&model)
	if err != nil {
		return nil, err
	}
	return model, nil
}
