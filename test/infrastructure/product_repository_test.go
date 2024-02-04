package infrastructure

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
