package main

import (
	"fmt"
	"github.com/devops787/go-docker-ci/router"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Handler: router.NewRouter(),
		Addr:    ":3000",
	}

	fmt.Println("Starting server at port 3000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
