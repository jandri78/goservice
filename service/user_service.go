package service

import (
	"github.com/jandri78/goservice/client"
	"github.com/jandri78/goservice/model"
)

func GetAllUsers() ([]model.User, error) {
	return client.FetchUsers()
}
