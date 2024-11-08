package main

import (
	"chat_app_server/db"
	"chat_app_server/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	r := mux.NewRouter()
	routes.InitRoutes(r)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:4321"})
	credentials := handlers.AllowCredentials()

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins, credentials)(r))
}
