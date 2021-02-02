package brand

import (
	"Clean_go/models"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreSt_GetByIdGetById(t *testing.T) {
	d,mock,err:=sqlmock.New()
	if err!=nil{
		fmt.Errorf("error while initializing mock")
	}
	db := New(d)
	testcases:= []models.Brands{
		{1,"LG"},
		{2,"Realme"},
	}
	for _,tc:=range testcases{
		query := `select id, name from brands where id=?`
		rows:= mock.NewRows([]string{"Id","Name"}).AddRow(tc.Id,tc.Name)
		mock.ExpectQuery(query).WithArgs(tc.Id).WillReturnRows(rows)
		res,_ := db.GetById(tc.Id)
		assert.Equal(t, tc,res)
	}
}

func TestGetByIdError(t *testing.T){
	d,mock,err:= sqlmock.New()
	if err!=nil{
		t.Error("error initialising mock")
	}
	db := New(d)
	testcases := []struct{
		id int
		err error
	}{
		{0,errors.New("invalid id")},
		{-1,errors.New("invalid id")},
	}
	for _,tc:= range testcases{
		query:= `select id, name from brands where id=?`
		mock.ExpectQuery(query).WithArgs(tc.id).WillReturnError(tc.err)
		_,err:= db.GetById(tc.id)
		assert.Equal(t, tc.err,err)
	}
}

func TestGetByName(t *testing.T) {
	d,mock,err := sqlmock.New()
	if err!=nil{
		t.Error("error initializing mock")
	}
	db := New(d)
	testcases:= []models.Brands{
		{1,"Television"},
		{2,"Realme"},
	}
	for _,tc:=range testcases{
		rows:= mock.NewRows([]string{"Id","Name"}).AddRow(tc.Id,tc.Name)
		query := `select id, name from brands where name=?`
		mock.ExpectQuery(query).WithArgs(tc.Name).WillReturnRows(rows)
		res,_:= db.GetByName(tc.Name)
		assert.Equal(t,tc,res)
	}
}


func TestGetByNameError(t *testing.T){
	d,mock,err:= sqlmock.New()
	if err!=nil{
		t.Error("error initialising mock")
	}
	db := New(d)
	testcases := []struct{
		name string
		err error
	}{
		{"LG",errors.New("invalid name")},
		{"Realme",errors.New("invalid name")},
	}
	for _,tc:= range testcases{
		query:= `select id, name from brands where name=?`
		mock.ExpectQuery(query).WithArgs(tc.name).WillReturnError(tc.err)
		_,err:= db.GetByName(tc.name)
		assert.Equal(t, tc.err,err)
	}
}


func TestCreate(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Error("error initializing sql mock")
	}
	db := New(d)
	testcases := []models.Brands{
		{1, "LG"},
		{2, "Realme"},
	}
	for i, tc := range testcases {
		//rows := mock.NewRows([]string{"Id", "Name", "Brand_ID"}).AddRow(tc.Id, tc.Name, tc.Brand_Id)
		query := `insert into brands`
		mock.ExpectExec(query).WithArgs(tc.Name).WillReturnResult(sqlmock.NewResult(int64(tc.Id),1))
		res,err:= db.Create(tc.Name)
		if err!=nil{
			t.Error("Failed!")
		}
		assert.Equal(t, testcases[i],res)
	}
}