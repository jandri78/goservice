package main

import (
	"log"
	"net/http"

	"github.com/jandri78/goservice/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Servidor en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
