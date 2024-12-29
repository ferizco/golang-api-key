package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// struct untuk menerima request dari client
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// struct untuk responsenya
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// handler untuk login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	apiKey := r.Header.Get("X-API-KEY")
	role, isValid := ValidateAPIKey(apiKey)

	if !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Status:  "error",
			Message: "API Key tidak valid",
		})
		return
	}
	log.Printf("api key valid, role: %s", role)

	if r.Method != http.MethodPost {
		http.Error(w, "Hanya method post yang diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  "error",
			Message: "Format JSON tidak valid",
		})
		return
	}

	if err := ValidateLogin(loginReq.Username, loginReq.Password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Status:  "success",
		Message: "Login valid",
	})
}
