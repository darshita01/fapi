package main

import (
	"backend/additionalFunc"
	"backend/business"
	"backend/database"
	"backend/handlers"
	"fmt"
	"log"
	"net/http"

	muxhan "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
) // //go:embed queryAddress.sql
// var SqlQueryAddress string

func main() {
	db := database.SetupDB()
	defer func() {
		err := db.Close()
		additionalFunc.CheckErr(err)
	}()
	router := mux.NewRouter()

	h := handlers.New(business.New(db))

	router.HandleFunc("/Users/{id}", h.GetData).Methods("GET")

	router.HandleFunc("/Users/", h.GetData).Methods("GET")

	router.HandleFunc("/Users/", h.EnterData).Methods("POST")

	router.HandleFunc("/Users/{id}", h.EditData).Methods("PUT")

	router.HandleFunc("/Users/{id}", h.DeleteData).Methods("Delete")

	fmt.Println("Server at ")
	//log.Fatal(http.ListenAndServe(":8000", muxhan.CORS()(router)))
	headersOk := muxhan.AllowedHeaders([]string{"*"})
	originsOk := muxhan.AllowedOrigins([]string{"*"})
	methodsOk := muxhan.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS","DELETE"})
	log.Fatal(http.ListenAndServe(":8000", muxhan.CORS(headersOk, originsOk, methodsOk)(router)))
}
