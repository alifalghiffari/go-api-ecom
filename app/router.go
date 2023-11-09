package app

import (
	"project-workshop/go-api-ecom/controller"
	"project-workshop/go-api-ecom/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController, 
			productController controller.ProductController,
			userController controller.UserController,
			accountController controller.AccountController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)

	router.POST("/api/accounts", accountController.UserDetailByID)

	router.PanicHandler = exception.ErrorHandler

	return router
}
