package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"product-app/domain"
)

var ctx context.Context

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	GetProductById(productId int64) (domain.Product, error)
	AddProduct(product domain.Product) error
	DeleteById(productId int64) error
	UpdatePrice(productId int64, price float32) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbpool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbpool,
	}
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx = context.Background()
	addProductSql := `Insert into products (name,price,discount,store) values ($1,$2,$3,$4)`

	addNewProduct, err := productRepository.dbPool.Exec(ctx, addProductSql, product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		log.Errorf("Error when adding product to table", err)
		return err
	}
	log.Info(fmt.Printf("Product added with %v", addNewProduct))
	return nil
}

func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx = context.Background()
	_, err := productRepository.GetProductById(productId)

	if err != nil {
		return errors.New("Product Not Found")
	}

	addProductSql := `Delete from products where id=$1`

	_, err = productRepository.dbPool.Exec(ctx, addProductSql, productId)

	if err != nil {
		log.Errorf("Error when deleting product ", err)
		return errors.New(fmt.Sprintf("Error while deleting product"))
	}
	log.Info(fmt.Printf("Product added with "))
	return nil
}

func (productRepository *ProductRepository) UpdatePrice(productId int64, price float32) error {
	ctx = context.Background()

	updateSql := `Update products set price = $1 where id=$2`

	_, err := productRepository.dbPool.Exec(ctx, updateSql, price, productId)

	if err != nil {
		return errors.New(fmt.Sprintf("Error while updating with id : %d", productId))
	}
	log.Info("Product %d price updated with new price %v", productId, price)
	return nil
}

func (productRepository *ProductRepository) GetProductById(productId int64) (domain.Product, error) {
	ctx = context.Background()
	getByIdSql := `Select * from products where id=$1`
	result := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)

	var product domain.Product
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	err := result.Scan(&id, &name, &price, &discount, &store)
	if err != nil {
		return domain.Product{}, errors.New(fmt.Sprintf("Error while getting product with id %v", productId))
	}
	product = domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}

	return product, nil

}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx = context.Background()
	getAllProductsSql := `Select * from products where store=$1`

	rows, err := productRepository.dbPool.Query(ctx, getAllProductsSql, storeName)
	if err != nil {
		log.Errorf("Error while getting products with store name")
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

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx = context.Background()
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
