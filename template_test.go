package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := "<html><body>{{.}}</body></html>"
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	requst := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, requst)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML")
}

func TestSimpleHTMLFile(t *testing.T) {
	requst := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, requst)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML")
}

func TestTemplateDirectory(t *testing.T) {
	requst := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, requst)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
