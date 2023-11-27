package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthMiddleware struct {
	Handler httprouter.Handle
}

func NewAuthMiddleware(handler httprouter.Handle) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Authentication logic or any other checks
	// Example: Fetch token from request header or cookie
	token := request.Header.Get("Authorization")
	if token == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte("Unauthorized"))
		return
	}

	// Pass the request to the next handler
	middleware.Handler(writer, request, nil)
}

func (middleware *AuthMiddleware) ApplyMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		middleware.ServeHTTP(writer, request)
	}
}
