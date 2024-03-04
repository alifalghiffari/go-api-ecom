package main

import (
	"fmt"
	"net/http"

	"project-workshop/go-api-ecom/app"
	"project-workshop/go-api-ecom/controller"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/repository"
	"project-workshop/go-api-ecom/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	accountRepository := repository.NewUserRepository()
	accountService := service.NewAccountService(accountRepository, db)
	accountController := controller.NewAccountController(accountService)

	cartRepository := repository.NewCartRepositoryImpl()
	cartService := service.NewCartService(cartRepository, userRepository, productRepository, db, validate)
	cartController := controller.NewCartController(cartService)

	orderRepository := repository.NewOrderRepositoryImpl()
	orderService := service.NewOrderService(orderRepository, userRepository, cartRepository, db, validate)
	orderController := controller.NewOrderController(orderService)

	router := app.NewRouter(categoryController, productController, accountController, userController, cartController, orderController)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	fmt.Println("Server listening on port http://localhost:3000/")

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
