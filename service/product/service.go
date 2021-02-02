package service

import (
	"Clean_go/errorConstant"
	"Clean_go/models"
	"Clean_go/service"
	"Clean_go/store"
	"errors"
	"strconv"
)

type serviceStruct struct {
	prod store.ProductStore
	brand store.BrandStore
}

func NewService(p store.ProductStore, b store.BrandStore) service.Service {
	return &serviceStruct{p,b}
}

func (s serviceStruct) GetProductDetails(id string) (models.Result,error){
	Id,_:= strconv.Atoi(id)
	//if er!=nil{
	//	return models.Result{},errors.New("invalid parameter id")
	//}
	prodRes,err := s.prod.GetById(Id)
	if err!=nil{
		return models.Result{},err
	}
	brandRes, er := s.brand.GetById(prodRes.Brand_Id)
	if er!=nil{
		return models.Result{},er
	}
	var ans models.Result
	ans.Id = prodRes.Id
	ans.ProductName = prodRes.Name
	ans.BrandName = brandRes.Name
	return ans,nil
}

func (s serviceStruct) CreateProduct (proName,brandName string) (models.Products,error){
	var pro models.Products
	brand, errr := s.brand.GetByName(brandName)

	// if brand does not exist
	if errr!=nil && errr.Error() == errorConstant.IdNotFound{
		var t models.Brands
		t.Name = brandName
		br,e := s.brand.Create(brandName)
		if e!=nil && br== (models.Brands{}){
			return models.Products{}, e
		}
		pro.Brand_Id = br.Id
	}else if errr!=nil{
		//brand not found
		return models.Products{}, errors.New(errorConstant.BrandIdNotFound)
	}
	pro.Name = proName
	if brand != (models.Brands{}) {
		pro.Brand_Id = brand.Id
	}

	// creating product entry
	ans,err := s.prod.Create(pro)
	if err!=nil {
		return models.Products{}, err
	}
	pro.Id =  ans.Id
	return pro,nil
}