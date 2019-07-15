package main

import (
	"net/http"
	"log"
	"fmt"
	"go-docker-ci/router"
)

func main() {
	server := http.Server{
		Handler: router.NewRouter(),
		Addr: ":3000",
	}

	fmt.Println("Starting server at port 3000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
