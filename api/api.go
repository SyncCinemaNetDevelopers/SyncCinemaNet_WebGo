package api

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

func getUsersList(w http.ResponseWriter, r *http.Request) {

}

func getUserInfo(w http.ResponseWriter, r *http.Request) {

}

func updateUserInfo(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func sendInvite(w http.ResponseWriter, r *http.Request){

}

/* rest api initialization
 * for more info read rest-api-spec.md
 */
func Initialize(r *mux.Router) {
	r.HandleFunc("/api/{ver}/server", getServersList)
	r.HandleFunc("/api/{ver}/server", addServerToList).Methods("POST")

	r.HandleFunc("/api/{ver}/server/{id}", getServerInfo).Methods("GET")
	r.HandleFunc("/api/{ver}/server/{id}", updateServerInfo).Methods("PUT")
	r.HandleFunc("/api/{ver}/server/{id}", deleteServer).Methods("DELETE")

	r.HandleFunc("/api/{ver}/room", getRoomsList).Methods("GET")
	r.HandleFunc("/api/{ver}/room", addRoomToList).Methods("POST")

	r.HandleFunc("/api/{ver}/room/{id}", getRoomInfo).Methods("GET")
	r.HandleFunc("/api/{ver}/room/{id}", updateRoomInfo).Methods("PUT")
	r.HandleFunc("/api/{ver}/room/{id}", deleteRoom).Methods("DELETE")

	r.HandleFunc("/api/{ver}/user",getUsersList).Methods("GET")

	r.HandleFunc("/api/{ver}/user/{nickname}",getUserInfo).Methods("GET")
	r.HandleFunc("/api/{ver}/user/{nickname}",updateUserInfo).Methods("PUT")
	r.HandleFunc("/api/{ver}/user/{nickname}",deleteUser).Methods("DELETE") // ??

	r.HandleFunc("/api/{ver}/user/{nickname}/send_invite/{room_id}",sendInvite).Methods("POST")
}
