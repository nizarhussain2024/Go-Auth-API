package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// In production, validate refresh token from database
	// For now, generate new tokens
	token := generateToken()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Token{
		AccessToken:  token,
		RefreshToken: generateToken(),
		TokenType:    "Bearer",
		ExpiresIn:    3600,
	})
}

func forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	store.mu.RLock()
	_, exists := store.users[req.Email]
	store.mu.RUnlock()

	// Always return success for security (don't reveal if email exists)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "If the email exists, a password reset link has been sent",
	})
}

