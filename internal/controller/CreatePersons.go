package controller

import (
	"fmt"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Criando uma nova pessoa\n")
	w.WriteHeader(http.StatusOK)
}
