package routes

import "github.com/gorilla/mux"

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/api/register", registerHandler).Methods("POST")
	r.HandleFunc("/api/login", loginHandler).Methods("POST")
	r.HandleFunc("/api/getPublicKey/{userid}", getPublicKeyHandler).Methods("GET")
	r.HandleFunc("/api/ws", relayHandler)
}
