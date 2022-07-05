package repository

import (
	"github.com/eminoz/customer-service-with-go/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}
type Order struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}

func UserCollection() *User {
	getDatabase := database.GetDatabase()
	return &User{
		Db:         getDatabase,
		Collection: getDatabase.Collection("user"),
	}
}

func OrderCollection() *Order {
	getDatabase := database.GetDatabase()
	return &Order{
		Db:         getDatabase,
		Collection: getDatabase.Collection("order"),
	}
}
