package main

import (
	"fmt"
	"net/http"

	"project-workshop/go-api-ecom/app"
	"project-workshop/go-api-ecom/controller"
	"project-workshop/go-api-ecom/helper"

	"project-workshop/go-api-ecom/middleware"
	"project-workshop/go-api-ecom/repository"
	"project-workshop/go-api-ecom/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
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

	router := app.NewRouter(categoryController, productController, userController, accountController)

	fmt.Println("Server listening on port http://localhost:3000/")

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
