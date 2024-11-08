package routes

import (
	"chat_app_server/db"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user db.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 16 {
		http.Error(w, "Password to short", http.StatusBadRequest)
		return
	}

	if user.PublicKey == "" {
		http.Error(w, "Public key not present", http.StatusBadRequest)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {

		http.Error(w, "Error while creating user", http.StatusBadRequest)
		return
	}
	user.Password = string(password)

	for {
		user.UserID = generateRandomString(20)
		err = db.DB.Create(&user).Error
		if err == nil {
			break
		}
	}

	cookie := generateCookie(user, w)
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(user.UserID)
}
