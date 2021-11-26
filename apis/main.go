package main

import (
	"apis/product_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	handler := cors.Default().Handler(router)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"X-Requested-With", "Content-TWype", "Authorization"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
	})
	handler = c.Handler(handler)
	router.HandleFunc("/api/product/findall", product_api.FindAll).Methods("GET")
	router.HandleFunc("/api/product/Create", product_api.Create).Methods("POST")
	router.HandleFunc("/api/product/Update", product_api.Update).Methods("PUT")
	router.HandleFunc("/api/product/Delete/{id}", product_api.Delete).Methods("DELETE")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		fmt.Println(err)
	}
}
