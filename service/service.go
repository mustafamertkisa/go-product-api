package service

import (
	"go-product-api/models"
	"go-product-api/repository"
)

type ProductService interface {
	GetAll() ([]models.Product, error)
	GetById(id string) (*models.Product, error)
	Create(req models.CreateProductRequest) (*models.Product, error)
	Update(id string, req models.UpdateProductRequest) error
	Delete(id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{repo: r}
}

func (s *productService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetById(id string) (*models.Product, error) {
	return s.repo.GetById(id)
}

func (s *productService) Create(req models.CreateProductRequest) (*models.Product, error) {
	product := models.Product{
		Name:  req.Name,
		Price: req.Price,
	}
	return s.repo.Create(product)
}

func (s *productService) Update(id string, req models.UpdateProductRequest) error {
	product := models.Product{
		Name:  req.Name,
		Price: req.Price,
	}
	return s.repo.Update(id, product)
}

func (s *productService) Delete(id string) error {
	return s.repo.Delete(id)
}
