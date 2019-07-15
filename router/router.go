package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"log"
	"github.com/devops787/go-docker-ci/calc"
	"encoding/json"
)

// Create HTTP router
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/add/{x:[0-9]+}/{y:[0-9]+}", AddHandler).Methods("GET")
	router.HandleFunc("/sub/{x:[0-9]+}/{y:[0-9]+}", SubtractHandler).Methods("GET")
	router.HandleFunc("/mul/{x:[0-9]+}/{y:[0-9]+}", MultiplyHandler).Methods("GET")
	router.HandleFunc("/div/{x:[0-9]+}/{y:[0-9]+}", DivideHandler).Methods("GET")

	return router
}

// Home route handler
func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Hello world!"))
}

// Add route handler
func AddHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	x, err := strconv.ParseInt(vars["x"], 10, 32)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	y, err := strconv.ParseInt(vars["y"], 10, 32)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result := calc.Add(int32(x), int32(y))
	response := map[string]interface{} {
		"result": result,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

// Subtract route handler
func SubtractHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	x, err := strconv.ParseInt(vars["x"], 10, 32)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	y, err := strconv.ParseInt(vars["y"], 10, 32)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result := calc.Subtract(int32(x), int32(y))
	response := map[string]interface{} {
		"result": result,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

// Multiply route handler
func MultiplyHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	x, err := strconv.ParseInt(vars["x"], 10, 32)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	y, err := strconv.ParseInt(vars["y"], 10, 32)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result := calc.Multiply(int32(x), int32(y))
	response := map[string]interface{} {
		"result": result,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

// Divide route handler
func DivideHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	x, err := strconv.ParseInt(vars["x"], 10, 32)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	y, err := strconv.ParseInt(vars["y"], 10, 32)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := calc.Divide(int32(x), int32(y))
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	response := map[string]interface{} {
		"result": result,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}