package route

import (
	"SMSM/pkg/utils"
	"log"
	"net/http"
)

func Launch() {
	http.HandleFunc("/", home)
	http.HandleFunc("/artistes", allArtistes)
	http.HandleFunc("/search", search)
	http.HandleFunc("/concert", concert)
	http.HandleFunc("/artiste/", artist)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	utils.Log("Server started at : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
