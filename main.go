package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"restapi"

	"fmt"
	"os"
)

func main() {
	r := mux.NewRouter()


	fmt.Fprintln(os.Stdout, http.ListenAndServe(":8080", r).Error())

}
