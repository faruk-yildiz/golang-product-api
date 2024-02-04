package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"product-app/domain"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	rows, err := productRepository.dbPool.Query(ctx, "Select * from products")
	if err != nil {
		log.Errorf("Error while getting all products %v", err)
		return []domain.Product{}
	}

	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for rows.Next() {
		err := rows.Scan(&id, &name, &price, &discount, &store)
		if err != nil {
			return nil
		}
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}
	return products
}
