package route

import (
	"net/http"
	"text/template"
)

type PageData struct {
	Title   string
	Content *template.Template
}

func home(w http.ResponseWriter, r *http.Request) {

}

func renderFile(title string, content *template.Template, w http.ResponseWriter) {
	data := PageData{
		Title:   title,
		Content: content,
	}

	t, _ := template.ParseFiles("template/layout.html")
	t.Execute(w, data)
}
