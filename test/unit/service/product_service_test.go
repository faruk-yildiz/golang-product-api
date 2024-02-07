package service

import (
	"github.com/stretchr/testify/assert"
	"os"
	"product-app/domain"
	"product-app/dto"
	"product-app/service"
	"testing"
)

var productService service.IProductService

func TestMain(t *testing.M) {

	initialProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    1000.0,
			Discount: 10,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 15,
			Store:    "ABC TECH",
		},
	}

	fakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductRepository)

	exitCode := t.Run()
	os.Exit(exitCode)
}

func TestGetAllProducts(t *testing.T) {
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))
	})
}
func TestWhenValidationErrorOccuredAddProduct(t *testing.T) {
	product := dto.ProductDto{
		Name:     "Test",
		Price:    1234,
		Discount: 74,
		Store:    "TEST STORE",
	}

	t.Run("GetAllProducts", func(t *testing.T) {
		err := productService.Add(product)
		assert.Equal(t, "Discount can not be greater than 70", err.Error())
	})
}
