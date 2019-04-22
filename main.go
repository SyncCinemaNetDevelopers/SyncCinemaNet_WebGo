/*Autor: Hamlet Avetikyn
 *Telegram: https://t.me/gregv2
 *GitHub: https://github.com/W1ll1am-Sh
 */
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	//hello
}

func logInHandler(w http.ResponseWriter, r *http.Request) {
	//login page
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//register page
}

func hotlistHandler(w http.ResponseWriter, r *http.Request) {
	//hotlist page
}

func privateHandler(w http.ResponseWriter, r *http.Request) {
	//private room page
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	//public room page
}

func createRoomHandler(w http.ResponseWriter, r *http.Request) {
	//page for creating room
}

func hereWasStone() boolean {
	return true
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/login", logInHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/hotlist", hotlistHandler)
	r.HandleFunc("/private", privateHandler)
	r.HandleFunc("/public", publicHandler)
	r.HandleFunc("/create_room", createRoomHandler)

	http.ListenAndServe(":8080", nil)
}
