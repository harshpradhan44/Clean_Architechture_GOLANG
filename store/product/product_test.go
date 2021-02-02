package product

import (
	"Clean_go/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestGetById(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Error("error initializing sqlmock")
	}
	db := New(d)
	testcases := []models.Products{
		{1, "Television", 1},
		{2, "Mobile", 3},
	}
	for i, tc := range testcases {
		rows := mock.NewRows([]string{"id", "name", "brand_id"}).AddRow(tc.Id, tc.Name, tc.Brand_Id)
		query := "select id,name,brand_Id from products where id=?"
		mock.ExpectQuery(query).WithArgs(tc.Id).WillReturnRows(rows)
		res,err:= db.GetById(tc.Id)
		if err!=nil{
			t.Error("Failed!")
		}
		if !assert.Equal(t, testcases[i],res){
			t.Errorf("failed at %v actual %v expected %v",i+1,res,testcases[i])
		}
	}
}

func TestGetByIdError(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Error("error initializing sqlmock")
	}
	db := New(d)
	testcases := []models.Products{
		{1, "Television", 1},
		{2, "Mobile", 3},
	}
	for i, tc := range testcases {
		rows := mock.NewRows([]string{"id", "name", "brand_id"}).AddRow(tc.Id, tc.Name, tc.Brand_Id)
		query := "select id,name,brand_Id from products where id=?"
		mock.ExpectQuery(query).WithArgs(tc.Id).WillReturnRows(rows)
		res,err:= db.GetById(tc.Id)
		if err!=nil{
			t.Error("Failed!")
		}
		if !assert.Equal(t, testcases[i],res){
			t.Errorf("failed at %v actual %v expected %v",i+1,res,testcases[i])
		}
	}
}

func TestGetByName(t *testing.T) {
	d,mock,err := sqlmock.New()
	if err!=nil{
		t.Error("error initializing mock")
	}
	db := New(d)
	testcases:= []models.Products{
		{1,"Television",2},
		//{2,"Realme"},
	}

	for _,tc:=range testcases{
		rows:= mock.NewRows([]string{"Id","Name","Brand_Id"}).AddRow(tc.Id,tc.Name,tc.Brand_Id)
		query := `select id,name,brand_Id from products where name =?`
		mock.ExpectQuery(query).WithArgs(tc.Name).WillReturnRows(rows)
		res,_:= db.GetByName(tc.Name)
		assert.Equal(t,tc,res)
	}
}

func TestCreate(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Error("error initializing sql mock")
	}
	db := New(d)
	testcases := []models.Products{
		{1, "Television", 1},
		{2, "Mobile", 3},
	}
	for i, tc := range testcases {
		//rows := mock.NewRows([]string{"Id", "Name", "Brand_ID"}).AddRow(tc.Id, tc.Name, tc.Brand_Id)
		query := "insert into products"
		mock.ExpectExec(query).WithArgs(tc.Id,tc.Name,tc.Brand_Id).WillReturnResult(sqlmock.NewResult(int64(tc.Id),1))
		res,err:= db.Create(tc)
		if err!=nil{
			t.Error("Failed!")
		}
		assert.Equal(t, testcases[i],res)
	}
}