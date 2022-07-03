package router

import (
	"github.com/eminoz/customer-service-with-go/api"
	"github.com/eminoz/customer-service-with-go/repository"
	"github.com/eminoz/customer-service-with-go/service"
	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	f := fiber.New()

	//User endpoints
	userCollection := repository.UserCollection()
	userEntity := repository.InsertOneEntity(userCollection) //giving userCollection struct
	userService := service.UserService{UserRepo: userEntity}
	controller := api.UserController{UserServices: &userService}
	u := f.Group("/user")
	u.Post("/createuser", controller.SaveUser)
	u.Get("/getOneUser/:id", controller.GetUserById)

	//Orders endpoints
	orderCollection := repository.OrderCollection()
	orderEntity := repository.InsertOneEntity(orderCollection) //giving orderCollection struct
	orderService := service.OrderService{OrderRepo: orderEntity}
	orderController := api.OrderController{OrderServices: &orderService}
	o := f.Group("/order")
	o.Post("/createorder", orderController.SaveOrder)
	o.Get("/getOneOrder/:id", orderController.GetOrderById)
	return f

}
