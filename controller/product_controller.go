package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"product-app/controller/request"
	"product-app/controller/response"
	"product-app/dto"
	"product-app/service"
	"strconv"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.GET("/api/v1/products", productController.GetAllProducts)
	e.POST("/api/v1/products", productController.CreateProduct)
	e.PUT("/api/v1/products/:id", productController.UpdateProduct)
	e.DELETE("/api/v1/products/:id", productController.DeleteProduct)
}

func (productContoller *ProductController) GetProductById(e echo.Context) error {
	param := e.Param("id")

	atoi, _ := strconv.Atoi(param)

	product, err := productContoller.productService.GetById(int64(atoi))

	if err != nil {
		return e.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, dto.ProductDto{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
}

func (productContoller *ProductController) GetAllProducts(e echo.Context) error {
	param := e.QueryParam("store")

	if len(param) == 0 {
		products := productContoller.productService.GetAllProducts()
		return e.JSON(http.StatusOK, products)
	}
	productsByStore := productContoller.productService.GetByStore(param)
	return e.JSON(http.StatusOK, productsByStore)
}
func (productContoller *ProductController) CreateProduct(e echo.Context) error {
	var addProductRequest request.AddProductRequest
	err := e.Bind(&addProductRequest)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}
	addProductErr := productContoller.productService.Add(addProductRequest.ToModel())
	if addProductErr != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: addProductErr.Error()})
	}
	return e.NoContent(http.StatusCreated)
}
func (productContoller *ProductController) UpdateProduct(e echo.Context) error {
	param := e.Param("id")
	atoi, _ := strconv.Atoi(param)
	productId := int64(atoi)

	newPrice := e.QueryParam("newPrice")
	if len(newPrice) == 0 {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{"New Price Required"})
	}

	priceFloat, err := strconv.ParseFloat(newPrice, 32)
	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{"New Price Format Wrong"})
	}
	productContoller.productService.UpdatePrice(productId, float32(priceFloat))
	return e.NoContent(http.StatusOK)
}
func (productContoller *ProductController) DeleteProduct(e echo.Context) error {
	param := e.Param("id")
	atoi, _ := strconv.Atoi(param)
	productId := int64(atoi)

	err := productContoller.productService.DeleteById(productId)

	if err != nil {
		return e.JSON(http.StatusNotFound, response.ErrorResponse{err.Error()})
	}
	return e.NoContent(http.StatusOK)
}
