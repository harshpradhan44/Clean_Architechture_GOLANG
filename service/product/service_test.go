package service

import (
	"Clean_go/errorConstant"
	"Clean_go/models"
	_ "Clean_go/service"
	"Clean_go/store"
	"errors"
	"github.com/golang/mock/gomock"
	_ "log"
	"reflect"
	"strconv"
	"testing"
)

func TestCreate(t *testing.T){
	ctrl := gomock.NewController(t)
	prod := store.NewMockProductStore(ctrl)
	brand := store.NewMockBrandStore(ctrl)
	prdsrv := NewService(prod,brand)
	pr := []models.Products{
		{0, "Mobile", 1},
		{0,"Mobile",3},
	}
	br := []models.Brands{
		{1,"Realme"},
		{3,"Sony"},
	}
	//ans := []models.Result{
	//	{1,"Mobile","Realme"},
	//	{2,"Television","LG"},
	//}
	testcases:= []struct{
		proName string
		brandName string
		prod models.Products
		brand models.Brands
		expected models.Products
		isBrandAvailable bool
		proError error
		brandError error
	}{
		//{"Mobile","Realme",pr[0],br[0],pr[0],true,nil,nil},
		{"Mobile","Sony",pr[1],br[1],pr[1],false,nil,errors.New(errorConstant.IdNotFound)},
	}

	for i,tc:=range testcases{

		if tc.isBrandAvailable {
			brand.EXPECT().GetByName(tc.brandName).Return(tc.brand, tc.brandError)
		}else{
			brand.EXPECT().GetByName(tc.brandName).Return(models.Brands{}, tc.brandError)
			brand.EXPECT().Create(tc.brandName).Return(tc.brand,tc.brandError)
		}

		prod.EXPECT().Create(tc.prod).Return(tc.expected,tc.proError)
		ans, _ := prdsrv.CreateProduct(tc.proName,tc.brandName)
		//if err!=tc.proError{
		//	t.Errorf("failed at %v got: %v actual: %v",i+1,err,tc.brandError)
		//}
		//if err!=tc.proError{
		//	t.Errorf("failed at %v got: %v actual: %v",i+1,err,tc.proError)
		//}
		if !reflect.DeepEqual(ans,tc.expected){
			//log.Fatal("Failed!")
			t.Errorf("failed at %v got: %v actual: %v",i+1,ans,tc.expected)
		}
	}
}

func TestGetById(t *testing.T){
	ctrl := gomock.NewController(t)
	prod := store.NewMockProductStore(ctrl)
	brand := store.NewMockBrandStore(ctrl)
	prdsrv := NewService(prod,brand)
	pr := []models.Products{
		{1,"Mobile",1},
		{2,"Television",4},
	}
	br := []models.Brands{
		{1,"Realme"},
		//{3,"LG"},
	}

	ans := []models.Result{
		{1,"Mobile","Realme"},
		{2,"Television","LG"},
	}
	testcases:= []struct{
		id int
		prod models.Products
		brand models.Brands
		expected models.Result
		proError error
		brandError error
	}{
		{1,pr[0],br[0],ans[0],nil,nil},
		{2,pr[1],models.Brands{},models.Result{},nil,errors.New(errorConstant.BrandIdNotFound)},
		{-1,models.Products{},models.Brands{},models.Result{},errors.New(errorConstant.IdNotFound),nil},
	}
	for i,tc:=range testcases{
		prod.EXPECT().GetById(tc.id).Return(tc.prod,tc.proError)
		if tc.prod!=(models.Products{}){
			brand.EXPECT().GetById(tc.prod.Brand_Id).Return(tc.brand,tc.brandError)
		}
		Id:= strconv.Itoa(tc.id)
		ans, err := prdsrv.GetProductDetails(Id)
		if err!=tc.brandError{
			t.Errorf("failed at %v got: %v actual: %v",i+1,err,tc.brandError)
		}
		if err!=tc.proError{
			t.Errorf("failed at %v got: %v actual: %v",i+1,err,tc.proError)
		}
		if ans != tc.expected {
			//log.Fatal("Failed!")
			t.Errorf("failed at %v got: %v actual: %v",i+1,ans,tc.expected)
		}
	}
}



