package service

import (
	"go.mongodb.org/mongo-driver/bson"
)

func makeFilter[filter any](field string, filterName filter) bson.D {
	d := bson.D{{field, filterName}}
	return d
}
