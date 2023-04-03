package controller

import (
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Criando um novo produto\n")
	w.WriteHeader(http.StatusOK)
}
