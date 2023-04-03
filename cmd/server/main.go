package main

import (
	"back-didge.com/back-end/internal/controller"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/createperson", controller.CreatePerson)
	http.HandleFunc("/createproduct", controller.CreateProduct)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
