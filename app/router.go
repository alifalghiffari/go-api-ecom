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
	userController controller.UserController,
	cartController controller.CartController,
	orderController controller.OrderController) *httprouter.Router {
	router := httprouter.New()

	// Middleware
	authMiddleware := middleware.Middleware{}

	// Order
	router.GET("/api/orders/users", authMiddleware.ApplyMiddleware(orderController.FindOrderByUserId))
	router.GET("/api/orders/edit/:orderId", authMiddleware.ApplyMiddleware(orderController.FindOrderById))
	router.POST("/api/orders", authMiddleware.ApplyMiddleware(orderController.CreateOrder))
	router.PUT("/api/orders/:orderId", authMiddleware.ApplyAdminMiddleware(orderController.UpdateOrder))

	// Keranjang
	router.POST("/api/carts", authMiddleware.ApplyMiddleware(cartController.AddToCart))
	router.PUT("/api/carts/:cartId", authMiddleware.ApplyMiddleware(cartController.UpdateCart))
	router.DELETE("/api/carts/:cartId", authMiddleware.ApplyMiddleware(cartController.DeleteCart))
	router.GET("/api/carts/edit/:cartId", authMiddleware.ApplyMiddleware(cartController.FindById))
	router.GET("/api/carts/users", authMiddleware.ApplyMiddleware(cartController.FindByUserId))
	router.GET("/api/carts", authMiddleware.ApplyMiddleware(cartController.FindAll))

	// Kategori
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", authMiddleware.ApplyAdminMiddleware(categoryController.Create))
	router.PUT("/api/categories/:categoryId", authMiddleware.ApplyAdminMiddleware(categoryController.Update))
	router.DELETE("/api/categories/:categoryId", authMiddleware.ApplyAdminMiddleware(categoryController.Delete))

	// Produk
	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", authMiddleware.ApplyAdminMiddleware(productController.Create))
	router.PUT("/api/products/:productId", authMiddleware.ApplyAdminMiddleware(productController.Update))
	router.DELETE("/api/products/:productId", authMiddleware.ApplyAdminMiddleware(productController.Delete))

	// Akun
	router.GET("/api/accounts", authMiddleware.ApplyMiddleware(accountController.UserDetailByID))

	// Pengguna
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}
