package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

// LoginHandler - POST /login (returns JWT token)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Load .env
	_ = godotenv.Load()

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Hardcoded login for demo – can replace with DB later
	if creds.Username != "Beast" || creds.Password != "Channu@4321" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		http.Error(w, "JWT_SECRET not set", http.StatusInternalServerError)
		return
	}

	// Token expiration time
	expHours := 2
	if v := os.Getenv("JWT_EXPIRE_HOURS"); v != "" {
		if h, err := strconv.Atoi(v); err == nil {
			expHours = h
		}
	}

	// Define JWT claims
	claims := jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Duration(expHours) * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		http.Error(w, "Token creation failed", http.StatusInternalServerError)
		return
	}

	// Return the token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TokenResponse{Token: tokenString})
}
