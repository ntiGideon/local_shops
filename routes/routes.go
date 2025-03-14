package routes

import (
	"baseProject/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController *controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/user/register", userController.Register)
	return router
}
