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

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	UserRepository  repository.UserRepository
	CartRepository  repository.CartRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, userRepository repository.UserRepository, cartRepository repository.CartRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		UserRepository:  userRepository,
		CartRepository:  cartRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func mapToOrderStatus(orderStatus string) domain.OrderStatus {
	switch orderStatus {
	case "PENDING":
		return domain.Pending
	case "PROCESS":
		return domain.Processing
	case "SHIPPING":
		return domain.Shipped
	case "DELIVERED":
		return domain.Delivered
	case "CANCELLED":
		return domain.Cancelled
	default:
		return domain.Pending
	}
}

func mapToPaymentStatus(paymentStatus string) domain.PaymentStatus {
	switch paymentStatus {
	case "PENDING":
		return domain.PaymentPending
	case "SUCCESS":
		return domain.PaymentSuccess
	case "FAILED":
		return domain.PaymentFailed
	default:
		return domain.PaymentPending
	}
}

func (service *OrderServiceImpl) CreateOrder(ctx context.Context, request web.OrderCreateRequest, userId int) web.OrderResponse {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user, err := service.UserRepository.FindById(ctx, tx, userId)
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }

    // Fetch cart items for the user
	cartItems, err := service.CartRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

    // Calculate total items and total price from cart items
    var totalItems int
    var totalPrice int

    for _, cartItem := range cartItems {
        // Iterate through the products in the cart item
        for _, product := range cartItem.Product {
            totalItems += cartItem.Quantity
            totalPrice += product.Price * cartItem.Quantity
        }

        // Delete the cart item
        service.CartRepository.DeleteCart(ctx, tx, cartItem)
    }

    // You can use the calculated totalItems and totalPrice to create the order
    orderStatus := mapToOrderStatus("PENDING")     // Set the default order status
    paymentStatus := mapToPaymentStatus("PENDING") // Set the default payment status

    order := domain.Order{
        UserID:        user.Id,
        CartId:        request.CartId,
        OrderItems:    cartItems,
        TotalItems:    totalItems,
        TotalPrice:    totalPrice,
        OrderStatus:   orderStatus,
        PaymentStatus: paymentStatus,
    }

    // Insert the order into the database
    order = service.OrderRepository.Insert(ctx, tx, order)

    return helper.ToOrderResponse(order)
}


func (service *OrderServiceImpl) UpdateOrder(ctx context.Context, request web.OrderUpdateRequest, Id int, userId int) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindById(ctx, tx, Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if order.UserID != userId {
		helper.PanicIfError(err)
	}

	orderStatus := mapToOrderStatus(request.OrderStatus)
	paymentStatus := mapToPaymentStatus(request.PaymentStatus)

	order = domain.Order{
		ID:            order.ID,
		UserID:        order.UserID,
		OrderStatus:   orderStatus,
		PaymentStatus: paymentStatus,
	}

	order = service.OrderRepository.Update(ctx, tx, order)

	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) FindOrderByUserId(ctx context.Context, userId int) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderResponses(order)
}

func (service *OrderServiceImpl) FindOrderById(ctx context.Context, Id int, userId int) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindById(ctx, tx, Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if order.UserID != userId {
		helper.PanicIfError(err)
	}

	return helper.ToOrderResponse(order)
}
