package product_api

import (
	"apis/config"
	"apis/entities"
	"apis/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	//enableCors(&response)
	db, err := config.GetMySQLDB()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, products)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	//enableCors(&response)
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetMySQLDB()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, product)
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	//enableCors(&response)
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	db, err := config.GetMySQLDB()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Update(&product)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, products)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	//enableCors(&response)
	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)

	db, err := config.GetMySQLDB()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		_, err2 := productModel.Delete(id)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, nil)
		}
	}
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	responseWithJson(w, code, map[string]string{"error": msg})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
