package route

import (
	"SMSM/pkg/utils"
	"log"
	"net/http"
)

// Launch initialise les routes HTTP de l'application et démarre le serveur web.
// Il configure les handlers pour les différentes pages et sert les fichiers statiques.
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
