package service

import "Clean_go/models"

type Service interface {
	GetProductDetails(id string) (models.Result,error)
	CreateProduct(productname, brandname string) (models.Products,error)
	//GetBrand(w http.ResponseWriter, r http.Request) (models.Brands,error)
}
