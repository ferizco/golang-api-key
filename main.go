package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", LoginHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Sesuaikan dengan origin frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	port := ":8000"
	log.Printf("Server berjalan di http://0.0.0.0%s\n", port)
	err := http.ListenAndServe("0.0.0.0"+port, corsHandler.Handler(mux))
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
