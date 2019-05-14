package restapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getServersList(w http.ResponseWriter, r *http.Request) {

}

func addServerToList(w http.ResponseWriter, r *http.Request) {

}

func getServerInfo(w http.ResponseWriter, r *http.Request) {

}

func updateServerInfo(w http.ResponseWriter, r *http.Request) {

}

func deleteServer(w http.ResponseWriter, r *http.Request) {

}

func getRoomsList(w http.ResponseWriter, r *http.Request) {

}

func addRoomToList(w http.ResponseWriter, r *http.Request) {

}

func getRoomInfo(w http.ResponseWriter, r *http.Request) {

}

func updateRoomInfo(w http.ResponseWriter, r *http.Request) {

}

func deleteRoom(w http.ResponseWriter, r *http.Request) {

}

func restAPIInitialize(r *mux.Router) {
	r.Handle("/api/{ver}/servers", getServersList).Methods("GET")
	r.Handle("/api/{ver}/servers", addServerToList).Methods("POST")

	r.Handle("/api/{ver}/servers/{id}", getServerInfo).Methods("GET")
	r.Handle("/api/{ver}/servers/{id}", updateServerInfo).Methods("PUT")
	r.Handle("/api/{ver}/servers/{id}", deleteServer).Methods("DELETE")

	r.Handle("/api/{ver}/rooms", getRoomsList).Methods("GET")
	r.Handle("/api/{ver}/rooms", addRoomToList).Methods("POST")

	r.Handle("/api/{ver}/rooms/{id}", getRoomInfo).Methods("GET")
	r.Handle("/api/{ver}/rooms/{id}", updateRoomInfo).Methods("PUT")
	r.Handle("/api/{ver}/rooms/{id}", deleteRoom).Methods("DELETE")
}