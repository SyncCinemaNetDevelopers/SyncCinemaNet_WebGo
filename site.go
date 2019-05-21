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
	w.Write([]byte("main page"))
}

func logInHandler(w http.ResponseWriter, r *http.Request) {
	//login page
	w.Write([]byte("LogIn page"))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//register page
	w.Write([]byte("SignUp page"))
}

func hotlistHandler(w http.ResponseWriter, r *http.Request) {
	//hotlist page
	w.Write([]byte("hotlist page"))
}

func privateHandler(w http.ResponseWriter, r *http.Request) {
	//private room page
	w.Write([]byte("private room page"))
}

func publicHandler(w http.ResponseWriter, r *http.Request) {
	//public room page
	w.Write([]byte("public room page"))
}

func createRoomHandler(w http.ResponseWriter, r *http.Request) {
	//page for creating room
	w.Write([]byte("creating room page"))
}

func siteInitialize(r *mux.Router) {
	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/login", logInHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/hotlist", hotlistHandler)
	r.HandleFunc("/private", privateHandler)
	r.HandleFunc("/public", publicHandler)
	r.HandleFunc("/create_room", createRoomHandler)
}
