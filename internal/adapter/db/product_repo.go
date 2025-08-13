package bd

import (
	"github.com/gomes800/go-products-api/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]domain.Product, error)
	FindByID(id uint) (domain.Product, error)
	Create(product domain.Product) (domain.Product, error)
	Update(product domain.Product) (domain.Product, error)
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func newProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	err := p.db.Find(&products).Error
	return products, err
}

func (p *productRepository) FindByID(id uint) (domain.Product, error) {
	var product domain.Product
	err := p.db.First(&product, id).Error
	return product, err
}

func (p *productRepository) Create(product domain.Product) (domain.Product, error) {
	err := p.db.Create(&product).Error
	return product, err
}

func (p *productRepository) Update(product domain.Product) (domain.Product, error) {
	err := p.db.Save(&product).Error
	return product, err
}

func (p *productRepository) Delete(id uint) error {
	return p.db.Delete(&domain.Product{}, id).Error
}
