package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"product-app/common/app"
	"product-app/common/postgre"
	"product-app/controller"
	"product-app/repository"
	"product-app/service"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	manager := app.NewConfigurationManager()

	dbPool := postgre.GetConnectionPool(ctx, manager.PostgreConfig)

	productRepository := repository.NewProductRepository(dbPool)

	productService := service.NewProductService(productRepository)

	productContoller := controller.NewProductController(productService)

	productContoller.RegisterRoutes(e)
	e.Start("localhost:8080")
}
