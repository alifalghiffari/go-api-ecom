package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CartController interface {
	AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}