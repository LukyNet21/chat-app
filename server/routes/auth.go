package routes

import (
	"chat_app_server/db"
	"crypto/rand"
	"math/big"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret string = "super-secret-key"

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

func generateRandomString(n int) string {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}

func generateCookie(user db.User, w http.ResponseWriter) *http.Cookie {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"iss": "chat-app",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now(),
	})

	token, err := claims.SignedString([]byte(JWTSecret))
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return nil
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HttpOnly: true,
		SameSite: 3, //1 - None, 2 - Lax, 3 - Strict
		Secure:   true,
	}
	return &cookie
}

func getUserFromCookie(cookie http.Cookie, w http.ResponseWriter) *db.User {
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		http.Error(w, "Could not parse token", http.StatusInternalServerError)
		return nil
	}

	if !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		var user db.User
		db.DB.Where("userid = ?", claims["sub"]).First(&user)

		return &user

	} else {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return nil
	}
}
