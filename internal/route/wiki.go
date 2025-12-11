package route

import (
	"SMSM/pkg/utils"
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

type PageData struct {
	TabTitle string
	Body     string
}

func home(w http.ResponseWriter, r *http.Request) {
	utils.Log("Page Home ouverte")
	data := HomeData{}

	renderFile("Accueil", render("home", data), w)
}

func renderFile(title string, content string, w http.ResponseWriter) {
	data := PageData{
		TabTitle: title,
		Body:     content,
	}

	executeHTML("layout", data, w)
}

func executeHTML(fileName string, data any, w http.ResponseWriter) {
	pathString := fmt.Sprintf("template/%v.html", fileName)

	t, _ := template.ParseFiles(pathString)
	t.Execute(w, data)
}

func render(fileName string, data any) string {
	pathString := fmt.Sprintf("template/%v.html", fileName)

	t, _ := template.ParseFiles(pathString)
	var buf bytes.Buffer
	t.Execute(&buf, data)
	return buf.String()
}
