package routes

import (
	"chat_app_server/db"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getPublicKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]

	var user db.User
	err := db.DB.Where("user_id = ?", userid).First(&user).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user.PublicKey)
}
