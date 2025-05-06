package client

import (
	"encoding/json"
	"net/http"

	"github.com/jandri78/goservice/model"
)

func FetchUsers() ([]model.User, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var users []model.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	return users, err
}
