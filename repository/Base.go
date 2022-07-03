package repository

import (
	"github.com/eminoz/customer-service-with-go/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type GenericCollection struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}

func UserCollection() *GenericCollection {
	getDatabase := database.GetDatabase()
	return &GenericCollection{
		Db:         getDatabase,
		Collection: getDatabase.Collection("user"),
	}
}

func OrderCollection() *GenericCollection {
	getDatabase := database.GetDatabase()
	return &GenericCollection{
		Db:         getDatabase,
		Collection: getDatabase.Collection("order"),
	}
}
