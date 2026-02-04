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
	var err error
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
	authURL := "https://accounts.spotify.com/api/token"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", authURL, strings.NewReader(data.Encode()))
	if err != nil {
		utils.LogError("", err)
	}

	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		utils.LogError("", err)
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
		utils.LogError("Décodage du token", err)
	}
	return tokenRes.AccessToken, nil
}

func GetTopArtist() {
	utils.Debug("Starting GetTopArtist function")
	// ID de la playlist Top 50 France
	playlistID := "2IgPkhcHbgQ4s4PdCxljAx"
	utils.Debug(fmt.Sprintf("DEBUG: Playlist ID set to: '%s'", playlistID))

	// Ajout du paramètre fields pour optimiser et market pour éviter les 404 sur les playlists officielles
	apiURL := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s?market=FR&fields=name,tracks.items(track(name,artists(name)))", playlistID)
	utils.Debug(fmt.Sprintf("DEBUG: Constructed API URL: %s", apiURL))

	utils.Debug(fmt.Sprintf("Fetching top artists from playlist %s...", playlistID))

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		utils.LogError("Erreur création requête", err)
		return
	}
	utils.Debug(fmt.Sprintf("DEBUG: Created HTTP GET request to: %s", apiURL))
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	utils.Debug("DEBUG: Added Authorization header with Bearer token")

	client := &http.Client{Timeout: 10 * time.Second}
	utils.Debug("DEBUG: Created HTTP client with 10s timeout")
	resp, err := client.Do(req)
	if err != nil {
		utils.LogError("Erreur requête HTTP", err)
		return
	}
	utils.Debug("DEBUG: HTTP request completed successfully")
	defer resp.Body.Close()
	utils.Debug("DEBUG: Response body will be closed on function exit")

	// Vérification du code HTTP
	utils.Debug(fmt.Sprintf("DEBUG: Response status code: %d\n", resp.StatusCode))
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		utils.Debug(fmt.Sprintf("DEBUG: Non-OK status code received: %d. Body: %s\n", resp.StatusCode, string(bodyBytes)))
		utils.Debug(fmt.Sprintf("Erreur API Spotify ! Code: %d\n", resp.StatusCode))
		return
	}
	utils.Debug("DEBUG: Status code is OK (200)")

	var result SpotifyPlaylist
	utils.Debug("DEBUG: Starting JSON decoding into SpotifyPlaylist struct")
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		utils.LogError("Erreur décodage JSON", err)
		return
	}
	utils.Debug("DEBUG: JSON decoding completed successfully")

	utils.Debug(fmt.Sprintf("Successfully decoded playlist: %s\n", result))
	utils.Debug(fmt.Sprintf("DEBUG: Playlist name: %s\n", result.Name))
	utils.Debug(fmt.Sprintf("DEBUG: Number of tracks: %d\n", len(result.Tracks.Items)))

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
	utils.Debug("DEBUG: End of GetTopArtist function")
}
