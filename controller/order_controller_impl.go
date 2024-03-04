package controller

import (
	"net/http"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) FindOrderByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := request.Context().Value("userId").(int)

	orderResponse := controller.OrderService.FindOrderByUserId(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) FindOrderById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := request.Context().Value("userId").(int)

	Id, err := strconv.Atoi(params.ByName("Id"))
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.FindOrderById(request.Context(), Id, userId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) CreateOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderCreateRequest := web.OrderCreateRequest{}
	userId := request.Context().Value("userId").(int)
	helper.ReadFromRequestBody(request, &orderCreateRequest)

	orderResponse := controller.OrderService.CreateOrder(request.Context(), orderCreateRequest, userId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) UpdateOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderUpdateRequest := web.OrderUpdateRequest{}
	userId := request.Context().Value("userId").(int)

	Id, err := strconv.Atoi(params.ByName("Id"))
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.UpdateOrder(request.Context(), orderUpdateRequest, Id, userId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}