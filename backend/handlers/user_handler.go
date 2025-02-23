package handlers

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/bmarquescost/expense-tracker/models"
	"github.com/bmarquescost/expense-tracker/repositories"
)

type UserHandler struct {
	Repo *repositories.UserRepository
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	if len(user.ID) == 0 {
		http.Error(w, "Bad request - Username must be non empty string", http.StatusBadRequest)
		return
	}

	log.Println("Inserting user: %s", user)
	if err := h.Repo.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to insert user: %s", user)
        return
	}

	log.Println("User registered: ", user)
	w.WriteHeader(http.StatusCreated)
}


func (h *UserHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Authentica user endpoint")
	var user models.User 
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	if len(user.ID) == 0 {
		http.Error(w, "Bad request - Username must be non empty string", http.StatusBadRequest)
		return
	}

	log.Println("Authenticating user: ", user)
	userExists, err := h.Repo.CheckForUser(&user)
	
	if !userExists {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to check for user user: %s", user)
        return
	}

	log.Println("User authenticated: ", user)
	w.WriteHeader(http.StatusOK)
}
