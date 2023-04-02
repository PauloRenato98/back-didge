package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/createNewProduct", createNewProduct)
	http.HandleFunc("/createNewtable", createNewtable)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createNewProduct(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Criando um novo produto\n")
	w.WriteHeader(http.StatusOK)
}

func createNewtable(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Criando uma nova mersa\n")
	w.WriteHeader(http.StatusOK)
}
