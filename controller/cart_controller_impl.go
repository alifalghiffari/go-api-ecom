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

func (controller *CartControllerImpl) AddToCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := web.CartCreateRequest{}
	helper.ReadFromRequestBody(r, &request)

	response := controller.CartService.AddToCart(r.Context(), request)

	helper.WriteToResponseBody(w, response)
}

func (controller *CartControllerImpl) UpdateCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := web.CartUpdateRequest{}
	helper.ReadFromRequestBody(r, &request)

	response := controller.CartService.UpdateCart(r.Context(), request)

	helper.WriteToResponseBody(w, response)
}

func (controller *CartControllerImpl) RemoveFromCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := web.CartRemoveRequest{}
	helper.ReadFromRequestBody(r, &request)

	controller.CartService.RemoveFromCart(r.Context(), request)

	w.WriteHeader(http.StatusOK)
}

func (controller *CartControllerImpl) GetItemsInCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId, _ := strconv.Atoi(r.URL.Query().Get("userId"))

	response := controller.CartService.GetItemsInCart(r.Context(), userId)

	helper.WriteToResponseBody(w, response)
}

func (controller *CartControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	response := controller.CartService.FindById(r.Context(), id)

	helper.WriteToResponseBody(w, response)
}