package main

import (
	"SyncCinemaNet_WebGo/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api.Initialize(r)
}
