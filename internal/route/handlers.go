package route

import "net/http"

func Launch() {
	http.HandleFunc("/", home)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
