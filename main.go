package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"fmt"
	"os"
)

func main() {
	r := mux.NewRouter()

	siteInitialize(r)

	fmt.Fprintln(os.Stdout, http.ListenAndServe(":8080", r).Error())

}
