package main

import (
	product3 "Clean_go/Handlers/product"
	product2 "Clean_go/service/product"
	"Clean_go/store/brand"
	"Clean_go/store/product"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	username = "root"
	host = "127.0.0.1"
	password = "Harsh@123"
	database = "harshDB"
)

func connectionString() string{
	return fmt.Sprintf("%v:%v@(%v)/%v",username,password,host,database)
}

func main(){
	fmt.Println("Server Running...")
	db, err := sql.Open("mysql", connectionString())
	if err!=nil{
		fmt.Println("error connecting database")
	}
	prod := product.New(db)
	brand := brand.New(db)
	ser := product2.NewService(prod,brand)
	handle := product3.Handler{Service: ser}
	router := mux.NewRouter()
	router.HandleFunc("/product",handle.Get).Methods("GET")
	router.HandleFunc("/product",handle.Create).Methods("POST")
	//fmt.Println(temp.GetById(3))
	http.ListenAndServe(":8080", router)
	defer db.Close()
}