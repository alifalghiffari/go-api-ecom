package app

import (
	"project-workshop/go-api-ecom/controller"
	"project-workshop/go-api-ecom/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController, productController controller.ProductController, usercontroller controller.UserController) *httprouter.Router {
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

	router.POST("/api/users", usercontroller.Create)
	router.PUT("/api/users/:userId", usercontroller.Update)
	router.DELETE("/api/users/:userId", usercontroller.Delete)
	router.GET("/api/users/:userId", usercontroller.FindById)

	router.PanicHandler = exception.ErrorHandler

	return router
}
