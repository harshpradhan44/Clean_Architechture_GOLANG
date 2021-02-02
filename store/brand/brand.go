package brand

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

func New(d *sql.DB) store.BrandStore {
	return storeSt{
		d,
	}
}

func (sql storeSt) GetById(id int) (models.Brands, error) {
	query := `select id, name from brands where id=?`
	res, er := sql.db.Query(query, id)

	if er != nil {
		fmt.Errorf("error executing query")
		return models.Brands{}, er
	}
	var temp models.Brands
	for res.Next(){
		err := res.Scan(&temp.Id, &temp.Name)
		if err!=nil{
			log.Fatal(err)
		}
	}
	// return result here
	return temp, nil
}

func (sql storeSt) GetByName(name string) (models.Brands, error) {
	query := `select id, name from brands where name=?`
	res, er := sql.db.Query(query,name)

	if er != nil {
		fmt.Errorf("error executing query")
		return models.Brands{}, er
	}
	var temp models.Brands
	for res.Next(){
		err := res.Scan(&temp.Id, &temp.Name)
		if err!=nil{
			fmt.Print("error in brand.go")
		}
	}
	if temp == (models.Brands{}){
		return models.Brands{}, errors.New("id not found")
	}
	// return result here
	return temp, nil
}

func (sql storeSt) Create (name string) (models.Brands,error){
	query := `insert into brands(name) values(?)`
	res,err := sql.db.Exec(query,name)
	if err!=nil{
		return models.Brands{}, err
	}
	var brand models.Brands
	rowAff,err := res.RowsAffected()
	if rowAff==0{
		return models.Brands{}, errors.New(errorConstant.IdAlreadyExist)
	}
	id,err := res.LastInsertId()
	brand.Name = name
	brand.Id = int(id)
	return brand, nil
}



