package router

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"strings"
)

// Home route handler
func TestHomeHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	http.HandlerFunc(HomeHandler).ServeHTTP(recorder, request)

	response := recorder.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", response.StatusCode)
	}
}

// Add route handler
func TestAddHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/add/1/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/add/{x:[0-9]+}/{y:[0-9]+}", AddHandler).Methods("GET")
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Invalid HTTP response")
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", response.StatusCode)
	}

	if response.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Unexpected content type: %s", response.Header.Get("Content-Type"))
	}

	result := strings.TrimSpace(string(body))
	expected := string(`{"result":3}`)

	if result != expected {
		t.Errorf("Invalid result: %s; expected: %s", result, expected)
	}
}

// Subtract route handler
func TestSubtractHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/sub/5/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/sub/{x:[0-9]+}/{y:[0-9]+}", SubtractHandler).Methods("GET")
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Invalid HTTP response")
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", response.StatusCode)
	}

	if response.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Unexpected content type: %s", response.Header.Get("Content-Type"))
	}

	result := strings.TrimSpace(string(body))
	expected := string(`{"result":3}`)

	if result != expected {
		t.Errorf("Invalid result: %s; expected: %s", result, expected)
	}
}

// Multiply route handler
func TestMultiplyHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/mul/5/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/mul/{x:[0-9]+}/{y:[0-9]+}", MultiplyHandler).Methods("GET")
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Invalid HTTP response")
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", response.StatusCode)
	}

	if response.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Unexpected content type: %s", response.Header.Get("Content-Type"))
	}

	result := strings.TrimSpace(string(body))
	expected := string(`{"result":10}`)

	if result != expected {
		t.Errorf("Invalid result: %s; expected: %s", result, expected)
	}
}

// Divide route handler
func TestDivideHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/div/8/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/div/{x:[0-9]+}/{y:[0-9]+}", DivideHandler).Methods("GET")
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Invalid HTTP response")
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", response.StatusCode)
	}

	if response.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Unexpected content type: %s", response.Header.Get("Content-Type"))
	}

	result := strings.TrimSpace(string(body))
	expected := string(`{"result":4}`)

	if result != expected {
		t.Errorf("Invalid result: %s; expected: %s", result, expected)
	}
}