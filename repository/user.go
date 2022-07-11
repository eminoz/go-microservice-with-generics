package repository

import (
	"github.com/eminoz/customer-service-with-go/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserModel model.User
var UserDal model.UserDal

func (u *User) InsertOne(ctx *fiber.Ctx, models interface{}) (interface{}, error) {
	userModel := models.(*model.User)

	emailFilter := makeFilter("email", userModel.Email) // make new bson filter

	isExist := u.IsExist(ctx, emailFilter) // check if user exist this email
	if isExist {
		return nil, &fiber.Error{Code: 409, Message: "Email already exist"}
	}

	insertOne, err := u.Collection.InsertOne(ctx.Context(), models) //insert new user
	if err != nil {
		return nil, err
	}
	filter := makeFilter("_id", insertOne.InsertedID)

	u.Collection.FindOne(ctx.Context(), filter).Decode(&UserDal) //find user after insert and decode that user userDal
	return UserDal, nil

}
func (u *User) GetOneByID(ctx *fiber.Ctx, id string) (interface{}, error) {
	userId, _ := primitive.ObjectIDFromHex(id)
	filter := makeFilter("_id", userId)

	err := u.Collection.FindOne(ctx.Context(), filter).Decode(&UserDal)
	if err != nil {
		return nil, err
	}
	return UserDal, nil
}

func (u *User) UpdateOneById(ctx *fiber.Ctx, id string, update interface{}) (interface{}, error) {

	userid, _ := primitive.ObjectIDFromHex(id)
	filter := makeFilter("_id", userid)
	isExist := u.IsExist(ctx, filter)
	if isExist == true {
		fil, up := makeFilterAndUpdate("_id", "$set", userid, update)

		updateOne, err := u.Collection.UpdateOne(ctx.Context(), fil, up)
		if err != nil {
			return nil, err
		}
		if updateOne.ModifiedCount == 1 {
			return "user updated", nil

		}
		return nil, &fiber.Error{Code: 409, Message: "user  already same "}
	}
	return nil, &fiber.Error{Code: 404, Message: "user  not found"}

}
func (u *User) GetAll(ctx *fiber.Ctx) (interface{}, error) {
	find, err := u.Collection.Find(ctx.Context(), bson.D{})
	if err != nil {
		return nil, err
	}
	var result []bson.M
	find.All(ctx.Context(), &result)
	return result, nil
}

func (u *User) IsExist(ctx *fiber.Ctx, filter bson.D) bool {
	documents, _ := u.Collection.CountDocuments(ctx.Context(), filter)
	if documents != 1 {
		return false
	}
	return true
}
