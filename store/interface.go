package store
import "Clean_go/models"

type ProductStore interface {
	GetById(id int) (models.Products,error)
	GetByName(name string) (models.Products,error)
	Create(products models.Products) (models.Products,error)
}

type BrandStore interface {
	GetById(id int) (models.Brands,error)
	GetByName (name string) (models.Brands,error)
	Create(name string) (models.Brands,error)
}
