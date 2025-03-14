package controller

import (
	"baseProject/dto"
	"baseProject/helpers"
	"baseProject/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: &userService}
}

func (controller *UserController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userDto := dto.RegisterUserDto{}
	helpers.ReadRequestBody(r, userDto)
	webResponse := controller.UserService.Register(r.Context(), &userDto)
	helpers.WriteResponseBody(w, webResponse, webResponse.Code)
}
