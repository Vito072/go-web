package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Vito"}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Fazrin",
	})
}

func TestTemplateFunction(t *testing.T) {
	requst := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, requst)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
