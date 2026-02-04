package api

import (
	"SMSM/internal/config"
	"SMSM/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	token       TokenResponse
	spotify_key string
	spotify_app string
)

func InitSpotify() {
	spotify_key, spotify_app = config.GetSpotifyKeyAndApp()
	var err error // Changé 'any' en 'error' pour être plus standard
	token.AccessToken, err = getAccessToken(spotify_app, spotify_key)
	if err != nil {
		utils.LogError("Erreur de la récupération du token", err)
		return
	}

	GetTopArtist()
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func getAccessToken(clientID, clientSecret string) (string, error) {
	// 1. Utiliser l'URL officielle d'authentification
	authURL := "https://accounts.spotify.com/api/token"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", authURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	// L'authentification Basic doit être faite avec le ClientID et ClientSecret
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// Lire le body pour voir le message d'erreur précis de Spotify
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		utils.LogError("Réponse Spotify: "+bodyString, nil)
		return "", fmt.Errorf("erreur auth spotify: code %d", resp.StatusCode)
	}

	var tokenRes TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenRes); err != nil {
		return "", err
	}
	return tokenRes.AccessToken, nil
}

func GetTopArtist() {
	fmt.Printf("DEBUG: Starting GetTopArtist function\n")
	playlistID := "37i9dQZEVXbKQ1ogMOyW9N"
	fmt.Printf("DEBUG: Playlist ID set to: %s\n", playlistID)

	// 2. Utiliser l'URL officielle de l'API v1
	apiURL := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", playlistID)
	fmt.Printf("DEBUG: Constructed API URL: %s\n", apiURL)

	fmt.Printf("Fetching top artists from playlist %s...\n", playlistID)

	req, _ := http.NewRequest("GET", apiURL, nil)
	fmt.Printf("DEBUG: Created HTTP GET request to: %s\n", apiURL)
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	fmt.Printf("DEBUG: Added Authorization header with Bearer token\n")

	client := &http.Client{Timeout: 10 * time.Second}
	fmt.Printf("DEBUG: Created HTTP client with 10s timeout\n")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("DEBUG: HTTP request failed with error: %v\n", err)
		utils.LogError("Erreur requête HTTP", err)
		return
	}
	fmt.Printf("DEBUG: HTTP request completed successfully\n")
	defer resp.Body.Close()
	fmt.Printf("DEBUG: Response body will be closed on function exit\n")

	// Vérification du code HTTP
	fmt.Printf("DEBUG: Response status code: %d\n", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("DEBUG: Non-OK status code received: %d\n", resp.StatusCode)
		fmt.Printf("Erreur API Spotify ! Code: %d\n", resp.StatusCode)
		return
	}
	fmt.Printf("DEBUG: Status code is OK (200)\n")

	var result SpotifyPlaylist
	fmt.Printf("DEBUG: Starting JSON decoding into SpotifyPlaylist struct\n")
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("DEBUG: JSON decoding failed with error: %v\n", err)
		utils.LogError("Erreur décodage JSON", err)
		return
	}
	fmt.Printf("DEBUG: JSON decoding completed successfully\n")

	fmt.Printf("Successfully decoded playlist: %s\n", result)
	fmt.Printf("DEBUG: Playlist name: %s\n", result.Name)
	fmt.Printf("DEBUG: Number of tracks: %d\n", len(result.Tracks.Items))

	// artistesUniques := make(map[string]bool)
	// fmt.Printf("DEBUG: Initialized map for unique artists\n")

	// // Notez le chemin d'accès : result -> Tracks -> Items
	// for i, item := range result.Tracks.Items {
	// 	fmt.Printf("DEBUG: Processing track item %d\n", i)
	// 	trackName := item.Track.Name
	// 	fmt.Printf("DEBUG: Track name: %s\n", trackName)
	// 	for j, artist := range item.Track.Artists {
	// 		fmt.Printf("DEBUG: Processing artist %d for track: %s\n", j, artist.Name)
	// 		// Affichage pour debug
	// 		// fmt.Printf("Artiste : %s (Titre : %s)\n", artist.Name, trackName)

	// 		if !artistesUniques[artist.Name] {
	// 			fmt.Printf("DEBUG: New unique artist found: %s\n", artist.Name)
	// 			artistesUniques[artist.Name] = true
	// 			// On peut afficher ici les artistes uniques trouvés
	// 			fmt.Println("- " + artist.Name + " (" + trackName + ")")
	// 		} else {
	// 			fmt.Printf("DEBUG: Artist %s already in unique list\n", artist.Name)
	// 		}
	// 	}
	// }

	// fmt.Printf("Total unique artists found: %d\n", len(artistesUniques))
	// fmt.Printf("DEBUG: Finished processing all tracks\n")
	fmt.Printf("DEBUG: End of GetTopArtist function\n")
}
