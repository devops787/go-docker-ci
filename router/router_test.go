package router

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"fmt"
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
	request, err := http.Get("/add/1/2")
	if err != nil {
		fmt.Println(err)

		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	http.HandlerFunc(AddHandler).ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
	if response.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", response.StatusCode)
	}
}
//
//// Subtract route handler
//func TestSubtractHandler(t *testing.T) {
//	request, err := http.NewRequest("GET", "/sub/2/2", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	recorder := httptest.NewRecorder()
//	http.HandlerFunc(SubtractHandler).ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	if response.StatusCode != http.StatusOK {
//		t.Errorf("Unexpected status code %d", response.StatusCode)
//	}
//}
//
//// Multiply route handler
//func TestMultiplyHandler(t *testing.T) {
//	request, err := http.NewRequest("GET", "/mul/2/2", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	recorder := httptest.NewRecorder()
//	http.HandlerFunc(MultiplyHandler).ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	if response.StatusCode != http.StatusOK {
//		t.Errorf("Unexpected status code %d", response.StatusCode)
//	}
//}
//
//// Divide route handler
//func TestDivideHandler(t *testing.T) {
//	request, err := http.NewRequest("GET", "/div/4/2", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	recorder := httptest.NewRecorder()
//	http.HandlerFunc(DivideHandler).ServeHTTP(recorder, request)
//
//	response := recorder.Result()
//	if response.StatusCode != http.StatusOK {
//		t.Errorf("Unexpected status code %d", response.StatusCode)
//	}
//}