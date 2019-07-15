package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"log"
	"go-docker-ci/calc"
	"encoding/json"
)

// Create HTTP router
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello world!"))
	})

	// add
	router.HandleFunc("/add/{x:[0-9]+}/{y:[0-9]+}", func(writer http.ResponseWriter, request *http.Request) {
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

		writer.WriteHeader(http.StatusOK)
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	}).Methods("GET")

	// subtract
	router.HandleFunc("/sub/{x:[0-9]+}/{y:[0-9]+}", func(writer http.ResponseWriter, request *http.Request) {
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

		writer.WriteHeader(http.StatusOK)
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	}).Methods("GET")

	// multiply
	router.HandleFunc("/mul/{x:[0-9]+}/{y:[0-9]+}", func(writer http.ResponseWriter, request *http.Request) {
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

		writer.WriteHeader(http.StatusOK)
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	}).Methods("GET")

	// divide
	router.HandleFunc("/div/{x:[0-9]+}/{y:[0-9]+}", func(writer http.ResponseWriter, request *http.Request) {
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

		writer.WriteHeader(http.StatusOK)
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	}).Methods("GET")

	return router
}
