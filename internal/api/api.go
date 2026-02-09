package api

import (
	"SMSM/pkg/utils"
	"encoding/json"
	"net/http"
)

// Launch initialise les connexions aux API externes et configure les routes API.
func Launch() {
	utils.Log("Chargement de l'API...")

	InitSpotify()

	// Enregistrement des handlers de l'API
	http.HandleFunc("/api/search", handleSearch)
	http.HandleFunc("/api/artist", handleArtist)
	http.HandleFunc("/api/artist/top-tracks", handleArtistTopTracks)

	utils.Log("API chargés avec succe")
}

// handleSearch gère les requêtes de recherche d'artistes.
// Il attend un paramètre de requête "q".
func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query is required", http.StatusBadRequest)
		return
	}

	result, err := Search(query)
	if err != nil {
		utils.LogError("API Search Error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// handleArtist gère les requêtes pour récupérer les détails d'un artiste.
// Il attend un paramètre de requête "id".
func handleArtist(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	artist, err := GetArtist(id)
	if err != nil {
		utils.LogError("API Artist Error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artist)
}

// handleArtistTopTracks gère les requêtes pour récupérer les titres populaires d'un artiste.
// Il attend un paramètre de requête "id".
func handleArtistTopTracks(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	tracks, err := GetArtistTopTrack(id)
	if err != nil {
		utils.LogError("API Artist Top Tracks Error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
}
