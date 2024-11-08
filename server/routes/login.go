package routes

import (
	"chat_app_server/db"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var loginUser db.User
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Could not decode data", http.StatusBadRequest)
		return
	}

	var user db.User
	err := db.DB.Where("userid = ?", loginUser.UserID).First(&user).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	cookie := generateCookie(user, w)
	http.SetCookie(w, cookie)
}
