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

func (middleware *AuthMiddleware) AllowOrigin(w http.ResponseWriter, req *http.Request) {
	// localhost:9000 origin mendapat ijin akses
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// allow cookie
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Ambil token dari cookie yang dikirim ketika request
	cookie, err := request.Cookie("token")
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte("Unauthorized"))
		return
	}
	// return bad request ketika field token tidak ada
	if cookie == nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad Request"))
		return
	}

	//middleware untuk autentikasi token
	token := request.Header.Get("Authorization")
	if token == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte("Unauthorized"))
		return
	}

	middleware.Handler(writer, request, nil) // Parameter `params` kosong untuk handler.
}

func (middleware *AuthMiddleware) ApplyMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		middleware.ServeHTTP(writer, request)
	}
}

func (middleware *AuthMiddleware) ApplyMiddlewareForOptions(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		middleware.AllowOrigin(writer, request)
	}
}
