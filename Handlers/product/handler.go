package product

import (
	"Clean_go/errorConstant"
	"Clean_go/models"
	"Clean_go/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct{
	Service service.Service
}

//func New(h product.Service) Handler {
//	return Handler{h}
//}


func (h Handler) Get(w http.ResponseWriter, r *http.Request){
	i := r.URL.Query()["id"]
	id, er := strconv.Atoi(i[0])
	if er!=nil{
		w.WriteHeader(404)
		w.Write([]byte(errorConstant.InvalidParameterId))
		return
	}
	Id := strconv.Itoa(id)
	res,err := h.Service.GetProductDetails(Id)
	if err!=nil{
		w.WriteHeader(404)
		w.Write([]byte(errorConstant.IdNotFound))
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h Handler) Create(w http.ResponseWriter, r* http.Request){
	var body models.Create
	json.NewDecoder(r.Body).Decode(&body)
	ans,err:= h.Service.CreateProduct(body.ProductName,body.BrandName)
	if err!=nil{
		fmt.Fprint(w,"error inserting")
		return
	}
	//fmt.Println("helo",ans)
	err = json.NewEncoder(w).Encode(&ans)
	if err!=nil {
		fmt.Printf("This is the error ->%v",err)
	}
}

