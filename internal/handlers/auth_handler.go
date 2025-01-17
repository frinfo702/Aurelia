package handlers

import (
	"Aurelia/internal/domain/models"
	"Aurelia/internal/domain/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type AuthHandler struct {
	authUC usecase.AuthUsecase
}

func NewAuthHandler(uc usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC: uc}
}

// POST api/auth/signup
func (h *AuthHandler) SignUpHander(w http.ResponseWriter, req *http.Request) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid reques"})
		return
	}
	if user.UserPassword == "" || user.UserEmail == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid reques"})
		return
	}

	err = h.authUC.SignUp(&user)
	if err != nil {
		log.Println("error signup", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "signup successful"})

}

// POST /api/auth/login
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err login", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid request"})
		return
	}

	token, err := h.authUC.Login(req.Email, req.Password)
	if err != nil {
		log.Println("err login", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid email or password"})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"token": token})

}
