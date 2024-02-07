package service

import (
	"errors"
	"product-app/domain"
	"product-app/repository"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) repository.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}
func (fakeRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	var products []domain.Product
	for i := 0; i < len(fakeRepository.products); i++ {
		if fakeRepository.products[i].Store == storeName {
			products = append(products, fakeRepository.products[i])
		}
	}
	return products
}
func (fakeRepository *FakeProductRepository) GetProductById(productId int64) (domain.Product, error) {
	var product domain.Product
	for i := 0; i < len(fakeRepository.products); i++ {
		if fakeRepository.products[i].Id == productId {
			product = fakeRepository.products[i]
		}
	}
	if product.Price > 0 {
		return product, nil
	}
	return domain.Product{}, errors.New("Product Not Found")
}
func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products) + 1),
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}
func (fakeRepository *FakeProductRepository) DeleteById(productId int64) error {
	return nil
}
func (fakeRepository *FakeProductRepository) UpdatePrice(productId int64, price float32) error {
	return nil
}
