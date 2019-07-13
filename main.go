package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world!"))
	})

	server := http.Server{
		Handler: router,
		Addr: ":3000",
	}

	fmt.Println("Listening on port 3000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}