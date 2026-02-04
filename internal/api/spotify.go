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

	// GetTopArtist()  <-- removed call
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

func GetTopArtist() (*SpotifyPlaylist, error) {
	utils.Debug("Starting GetTopArtist function")
	// ID de la playlist Top 50 France
	playlistID := "2IgPkhcHbgQ4s4PdCxljAx"
	utils.Debug(fmt.Sprintf("DEBUG: Playlist ID set to: '%s'", playlistID))

	// Ajout du paramètre fields pour optimiser et market pour éviter les 404 sur les playlists officielles
	apiURL := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s?market=FR&fields=name,tracks.items(track(name,artists(name),album(images(url))))", playlistID)
	utils.Debug(fmt.Sprintf("DEBUG: Constructed API URL: %s", apiURL))

	utils.Debug(fmt.Sprintf("Fetching top artists from playlist %s...", playlistID))

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		utils.LogError("Erreur création requête Playlist", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		utils.LogError("Erreur requête HTTP Playlist", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		utils.Debug(fmt.Sprintf("DEBUG: Non-OK status code received: %d. Body: %s\n", resp.StatusCode, string(bodyBytes)))
		return nil, fmt.Errorf("playlist api error: %d", resp.StatusCode)
	}

	var result SpotifyPlaylist
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		utils.LogError("Erreur décodage JSON", err)
		return nil, err
	}

	utils.Debug(fmt.Sprintf("Successfully decoded playlist: %s\n", result.Name))

	return &result, nil
}

func GetArtist(artisteID string) (*SpotifyArtist, error) {
	utils.Debug("Starting GetArtist function")
	utils.Debug(fmt.Sprintf("DEBUG: Artiste ID set to: '%s'", artisteID))

	// Ajout du paramètre fields pour optimiser et market pour éviter les 404 sur les playlists officielles
	apiURL := fmt.Sprintf("https://api.spotify.com/v1/artists/%s", artisteID)
	utils.Debug(fmt.Sprintf("DEBUG: Constructed API URL: %s", apiURL))

	utils.Debug(fmt.Sprintf("Fetching data from artiste %s...", artisteID))

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		utils.LogError("Erreur création requête Artiste", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		utils.LogError("Erreur requête HTTP Artiste", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		utils.Debug(fmt.Sprintf("DEBUG: Non-OK status code received: %d. Body: %s\n", resp.StatusCode, string(bodyBytes)))
		return nil, fmt.Errorf("artiste api error: %d", resp.StatusCode)
	}

	var result SpotifyArtist
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		utils.LogError("Erreur décodage JSON", err)
		return nil, err
	}

	utils.Debug(fmt.Sprintf("Successfully decoded artiste: %s\n", result.Name))

	return &result, nil
}

func GetArtistTopTrack(artisteID string) (*SpotifyArtistTopTracks, error) {
	utils.Debug("Starting GetArtistTopTrack function")

	// API URL pour récupérer les top tracks d'un artiste (market=FR est obligatoire ou recommandé)
	apiURL := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/top-tracks?market=FR", artisteID)
	utils.Debug(fmt.Sprintf("DEBUG: Constructed API URL: %s", apiURL))

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		utils.LogError("Erreur création requête Artiste Top Tracks", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		utils.LogError("Erreur requête HTTP Artiste Top Tracks", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		utils.Debug(fmt.Sprintf("DEBUG: Non-OK status code received: %d. Body: %s\n", resp.StatusCode, string(bodyBytes)))
		return nil, fmt.Errorf("artiste top tracks api error: %d", resp.StatusCode)
	}

	var result SpotifyArtistTopTracks
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		utils.LogError("Erreur décodage JSON Top Tracks", err)
		return nil, err
	}

	utils.Debug(fmt.Sprintf("Successfully decoded %d tracks for artist %s\n", len(result.Tracks), artisteID))

	return &result, nil
}
