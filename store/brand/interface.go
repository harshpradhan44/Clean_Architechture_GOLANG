package brand
import "Clean_go/models"

type Store interface {
	GetById(id int) (models.Brands,error)
	GetByName (name string) (models.Brands,error)
	Create(name string) (models.Brands,error)
}
