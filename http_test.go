package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest("GET", "http:localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, *request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
