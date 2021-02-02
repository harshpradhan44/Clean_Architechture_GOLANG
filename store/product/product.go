package product

import (
	"Clean_go/errorConstant"
	"Clean_go/models"
	"Clean_go/store"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type storeSt struct {
	db *sql.DB
}

func New(d *sql.DB) store.ProductStore {
	return storeSt{
		d,
	}
}

func (sql storeSt) GetById(id int) (models.Products, error) {
	query := `select id,name,brand_Id from products where id=?`
	res, er := sql.db.Query(query, id)

	if er != nil {
		fmt.Errorf("error executing query")
		return models.Products{}, er
	}

	if res==nil{
		return models.Products{}, errors.New("not found")
	}
	var temp models.Products
	for res.Next(){
		err := res.Scan(&temp.Id, &temp.Name, &temp.Brand_Id)
		if err!=nil{
			log.Fatal(err)
		}
	}
	// return result here
	return temp, nil
}

func (sql storeSt) GetByName(name string) (models.Products, error) {
	query := `select id,name,brand_Id from products where name =?`
	res, er := sql.db.Query(query, name)

	if er != nil {
		fmt.Errorf("error executing query")
		return models.Products{}, er
	}

	//if res==nil{
	//	return models.Products{}, errors.New("not found")
	//}
	var temp models.Products
	for res.Next(){
		err := res.Scan(&temp.Id, &temp.Name, &temp.Brand_Id)
		if err!=nil{
			log.Fatal(err)
		}
	}
	if temp == (models.Products{}){
		return models.Products{}, errors.New(errorConstant.IdNotFound)
	}
	// return result here
	return temp, nil
}

func (sql storeSt) Create(pro models.Products) (models.Products,error){
	query := "insert into products values(?,?,?)"
	res,err := sql.db.Exec(query,pro.Id,pro.Name,pro.Brand_Id)
	if err!=nil{
		return models.Products{}, err
	}
	rowAff,err := res.RowsAffected()
	if rowAff==0{
		return models.Products{}, errors.New(errorConstant.IdAlreadyExist)
	}
	id,err := res.LastInsertId()
	pro.Id = int(id)
	return pro, nil
}

