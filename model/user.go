package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name     string `validate:"required,omitempty"`
	Surname  string `json:"surname"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,omitempty,gte=7,lte=130"`
}
type Login struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `json:"email",bson:"email"`
	Password string             `json:"password"`
}
type UserDal struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Email   string `json:"email" bson:"email"`
}
type LoginDal struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email"`
	Token string             `json:"token" bson:"token"`
}
type Email struct {
	Email string `json:"email"`
}
