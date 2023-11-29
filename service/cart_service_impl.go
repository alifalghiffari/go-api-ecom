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

type CartServiceImpl struct {
	CartRepository repository.CartRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewCartService(cartRepository repository.CartRepository, DB *sql.DB, validate *validator.Validate) CartService {
	return &CartServiceImpl{
		CartRepository: cartRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *CartServiceImpl) AddToCart(ctx context.Context, request web.CartCreateRequest) web.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.Cart{
		UserId:    request.UserId,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
	}

	cart = service.CartRepository.AddToCart(ctx, tx, cart)

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) UpdateCart(ctx context.Context, request web.CartUpdateRequest) web.CartResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.Cart{
		Id:       request.Id,
		Quantity: request.Quantity,
	}

	cart = service.CartRepository.UpdateCart(ctx, tx, cart)

	return helper.ToCartResponse(cart)
}

func (service *CartServiceImpl) RemoveFromCart(ctx context.Context, request web.CartRemoveRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.Cart{
		UserId:    request.UserId,
		ProductId: request.ProductId,
	}

	service.CartRepository.RemoveFromCart(ctx, tx, cart)
}

func (service *CartServiceImpl) GetItemsInCart(ctx context.Context, userId int) []web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	carts := service.CartRepository.GetItemsInCart(ctx, tx, userId)

	return helper.ToCartResponses(carts)
}

func (service *CartServiceImpl) FindById(ctx context.Context, id int) web.CartResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := service.CartRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCartResponse(cart)
}