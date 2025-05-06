package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jandri78/goservice/service"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
