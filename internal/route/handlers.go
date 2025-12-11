package route

import (
	"SMSM/pkg/utils"
	"log"
	"net/http"
)

func Launch() {
	http.HandleFunc("/", home)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	utils.Log("Server started at : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
