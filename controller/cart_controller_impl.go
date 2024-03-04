package controller

import (
	"net/http"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CartControllerImpl struct {
	CartService service.CartService
}

func NewCartController(cartService service.CartService) CartController {
	return &CartControllerImpl{
		CartService: cartService,
	}
}

func (controller *CartControllerImpl) AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartCreateRequest := web.CartCreateRequest{}
	userId := request.Context().Value("userId").(int)
	helper.ReadFromRequestBody(request, &cartCreateRequest)

	cartResponse := controller.CartService.AddToCart(request.Context(), cartCreateRequest, userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImpl) UpdateCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartUpdateRequest := web.CartUpdateRequest{}
	userId := request.Context().Value("userId").(int)
	helper.ReadFromRequestBody(request, &cartUpdateRequest)

	cartResponse := controller.CartService.UpdateCart(request.Context(), cartUpdateRequest, userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImpl) DeleteCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartDeleteRequest := web.CartDeleteRequest{}
	userId := request.Context().Value("userId").(int)
	helper.ReadFromRequestBody(request, &cartDeleteRequest)

	cartResponse := controller.CartService.DeleteCart(request.Context(), cartDeleteRequest, userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartId := params.ByName("cartId")
	id, err := strconv.Atoi(cartId)
	helper.PanicIfError(err)

	cartResponse := controller.CartService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImpl) FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := request.Context().Value("userId").(int)

	cartResponse := controller.CartService.FindByUserId(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartResponse := controller.CartService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}