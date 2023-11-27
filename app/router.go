package app

import (
	// "net/http"
	"project-workshop/go-api-ecom/controller"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController,
	productController controller.ProductController,
	accountController controller.AccountController,
	userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	// Middleware
	authMiddleware := middleware.AuthMiddleware{}

	// Kategori
	router.GET("/api/categories", (categoryController.FindAll))
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", authMiddleware.ApplyMiddleware(categoryController.Create))
	router.PUT("/api/categories/:categoryId", authMiddleware.ApplyMiddleware(categoryController.Update))
	router.DELETE("/api/categories/:categoryId", authMiddleware.ApplyMiddleware(categoryController.Delete))

	// Produk
	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", authMiddleware.ApplyMiddleware(productController.Create))
	router.PUT("/api/products/:productId", authMiddleware.ApplyMiddleware(productController.Update))
	router.DELETE("/api/products/:productId", authMiddleware.ApplyMiddleware(productController.Delete))

	// Akun
	router.POST("/api/accounts", authMiddleware.ApplyMiddleware(accountController.UserDetailByID))

	// Pengguna
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}
