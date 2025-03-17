package service

import (
	"baseProject/data"
	"baseProject/dto"
	"baseProject/ent"
	"baseProject/ent/user"
	"baseProject/helpers"
	"context"
	"net/http"
)

type UserService struct {
	Db *ent.Client
}

func NewUserService(e *ent.Client) *UserService {
	return &UserService{Db: e}
}

func (p *UserService) Register(ctx context.Context, userDto *dto.RegisterUserDto) *data.WebResponse {
	validatorErrors := helpers.RequestValidators(userDto)
	if validatorErrors != nil {
		return &data.WebResponse{
			Code:    http.StatusBadRequest,
			Message: validatorErrors.Error(),
			Data:    nil,
		}
	}

	existingUser, err := p.Db.User.Query().Where(user.EmailEQ(userDto.Email)).Only(ctx)

	if existingUser != nil {
		return &data.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "User already exists",
			Data:    nil,
		}
	}
	hashedPassword := helpers.HashPassword(userDto.Password)
	_, err = p.Db.User.Create().SetName(userDto.Name).SetEmail(userDto.Email).SetPhone(userDto.Phone).SetPassword(hashedPassword).Save(ctx)
	if err != nil {
		return &data.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return &data.WebResponse{
		Code:    http.StatusCreated,
		Message: "User created",
		Data:    nil,
	}
}
