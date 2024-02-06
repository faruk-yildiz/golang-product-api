package service

import (
	"errors"
	"product-app/domain"
	"product-app/dto"
	"product-app/repository"
)

type IProductService interface {
	Add(productDto dto.ProductDto) error
	DeleteById(productId int64) error
	GetById(productId int64) ([]domain.Product, error)
	UpdatePrice(productId int64, price float32) error
	GetAllProducts() []domain.Product
	GetByStore(storeName string) []domain.Product
}

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func validateProductCreate(productDto dto.ProductDto) error {
	if productDto.Discount > 70 {
		errors.New("Discount can not be greater than 70")
	}
	return nil
}

func (pService ProductService) Add(productDto dto.ProductDto) error {
	err := validateProductCreate(productDto)
	if err != nil {
		return err
	}
	return pService.productRepository.AddProduct(domain.Product{
		Name:     productDto.Name,
		Price:    productDto.Price,
		Discount: productDto.Discount,
		Store:    productDto.Store,
	})
}

func (pService ProductService) DeleteById(productId int64) error {
	return pService.productRepository.DeleteById(productId)
}
func (pService ProductService) GetById(productId int64) ([]domain.Product, error) {
	return pService.GetById(productId)
}
func (pService ProductService) UpdatePrice(productId int64, price float32) error {
	return pService.UpdatePrice(productId, price)
}
func (pService ProductService) GetAllProducts() []domain.Product {
	return pService.GetAllProducts()
}
func (pService ProductService) GetByStore(storeName string) []domain.Product {
	return pService.productRepository.GetAllProductsByStore(storeName)
}
