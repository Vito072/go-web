package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/layout.gohtml",
		"./templates/footer.gohtml",
	))
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Vito",
	})
}

func TestTemplateLayout(t *testing.T) {
	requst := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, requst)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
