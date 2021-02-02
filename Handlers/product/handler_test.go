package product

import (
	"Clean_go/errorConstant"
	"Clean_go/models"
	"Clean_go/service"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	prod := service.NewMockService(ctrl)
	a := Handler{prod}
	result := []models.Result{
		{13,"zopNow","zopsmart"},
		{},
	}
	testcases := []struct {
		id       string
		response string
		error error
	}{
		{"13", "200",nil},
		{"200","404",errors.New(errorConstant.IdNotFound)},
	}

	for i, tc:= range testcases {
		req,err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/product?id=%v",tc.id),nil)
		if err!=nil{
			t.Error(err)
		}

		w := httptest.NewRecorder()
		prod.EXPECT().GetProductDetails(tc.id).Return(result[i],tc.error)
		a.Get(w, req)
		if !reflect.DeepEqual(fmt.Sprint(w.Code),tc.response){
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, tc.response)
		}
	}
}

func TestGetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	prod := service.NewMockService(ctrl)
	a := Handler{prod}
	testcases := []struct {
			id string
			body string
			response string
	}{
		{"1a", "invalid parameter id","404"},
		//{"200", errorConstant.IdNotFound,"404"},
	}

	for i, tc:= range testcases {
		req,err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/product?id=%v",tc.id),nil)
		if err!=nil{
			t.Error(err)
		}

		w := httptest.NewRecorder()
		a.Get(w, req)

		if !reflect.DeepEqual(w.Body.String(),tc.body){
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body, tc.body)
		}
		if !reflect.DeepEqual(fmt.Sprint(w.Code),tc.response){
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, tc.response)
		}
	}
}