package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderController interface {
	FindOrderByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindOrderById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateOrder(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}