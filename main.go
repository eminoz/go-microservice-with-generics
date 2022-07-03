package main

import (
	"github.com/eminoz/customer-service-with-go/pkg/config"
	"github.com/eminoz/customer-service-with-go/pkg/database"
	"github.com/eminoz/customer-service-with-go/pkg/router"
)

func main() {
	config.SetupConfig()
	database.SetDatabase()
	setup := router.Setup()
	setup.Listen(":" + config.GetConfig().Port)
}
