package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Adjust for better security
		return true
	},
}

func relayHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "Unauthorize", http.StatusUnauthorized)
		return
	}
	user := getUserFromCookie(*cookie, w)
	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		conn.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received message: %s", message)
	}
}
