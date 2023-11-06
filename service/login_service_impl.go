package service

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
)

type LoginServiceImpl struct {
	LoginRepository repository.LoginRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewLoginServiceImpl(loginRepository repository.LoginRepository, DB *sql.DB, validate *validator.Validate) LoginService {
	return &LoginServiceImpl{
		LoginRepository: loginRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *LoginServiceImpl) Login(ctx context.Context, tx *string, login web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(login)
	helper.PanicIfError(err)

	if tx != nil {
		panic(exception.NewUnauthorized("You are already logged in"))
	}

	txDB, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(txDB)

	user := service.LoginRepository.FindByUsername(ctx, txDB, login.Username)

	if user.Id == 0 {
		panic(exception.NewUnauthorized("Username or password is wrong"))
	}

	if user.Password != login.Password {
		panic(exception.NewUnauthorized("Username or password is wrong"))
	}

	token := helper.GenerateToken(user.Id, user.Role)

	return web.LoginResponse{
		Id:    user.Id,
		Token: token,
	}
}


