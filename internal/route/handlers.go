package route

import (
	"fmt"
	"log"
	"net/http"
)

func Launch() {
	http.HandleFunc("/", home)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server launch at : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
