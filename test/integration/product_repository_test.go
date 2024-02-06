package integration

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"os"
	"product-app/common/postgre"
	"product-app/domain"
	"product-app/repository"
	"testing"
)

var productRepository repository.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgre.GetConnectionPool(ctx, postgre.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	productRepository = repository.NewProductRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbpool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbpool)
}

func clear(ctx context.Context, dbpool *pgxpool.Pool) {
	TruncateTestData(ctx, dbpool)
}

func TestGetAllProducts(t *testing.T) {
	setup(ctx, dbPool)
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}

func TestAddProduct(t *testing.T) {
	ctx = context.Background()
	expectedProducts := []domain.Product{
		{
			Name:     "Test Urun",
			Price:    45,
			Discount: 10,
			Store:    "TEST STORE",
		},
	}

	product := domain.Product{
		Name:     "Test Urun",
		Price:    45,
		Discount: 10,
		Store:    "TEST STORE",
	}

	t.Run("AddProduct", func(t *testing.T) {
		productRepository.AddProduct(product)
		products := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(products))
		assert.Equal(t, expectedProducts, products)
	})
	clear(ctx, dbPool)
}

func TestUpdatePrice(t *testing.T) {
	setup(ctx, dbPool)
	t.Run("Update Price", func(t *testing.T) {
		productBeforeUpdate, _ := productRepository.GetProductById(1)
		assert.Equal(t, float32(3000.0), productBeforeUpdate.Price)
		productRepository.UpdatePrice(1, 4000.0)
		productAfterUpdate, _ := productRepository.GetProductById(1)
		assert.Equal(t, float32(4000.0), productAfterUpdate.Price)
	})
	clear(ctx, dbPool)
}

func TestGetAllProductsByStore(t *testing.T) {
	setup(ctx, dbPool)
	storeName := "ABC TECH"
	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore(storeName)
		assert.Equal(t, 3, len(actualProducts))
	})
	clear(ctx, dbPool)
}

func TestGetProductById(t *testing.T) {
	setup(ctx, dbPool)
	var id int64 = 2

	expectedProuduct := domain.Product{
		Id:       2,
		Name:     "Ütü",
		Price:    1500.0,
		Discount: 10.0,
		Store:    "ABC TECH",
	}

	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProuduct, _ := productRepository.GetProductById(id)
		assert.Equal(t, actualProuduct, expectedProuduct)
	})
	clear(ctx, dbPool)
}

func TestDeleteProduct(t *testing.T) {
	setup(ctx, dbPool)

	ctx = context.Background()
	var productId int64 = 1

	t.Run("AddProduct", func(t *testing.T) {
		productRepository.DeleteById(productId)
		products := productRepository.GetAllProducts()

		assert.Equal(t, 3, len(products))
	})
	clear(ctx, dbPool)
}
