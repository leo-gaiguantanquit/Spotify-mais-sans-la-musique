package api

import (
	"SMSM/internal/config"
	"SMSM/pkg/utils"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

var (
	token       TokenResponse
	spotify_key string
	spotify_app string
)

func initSpotifyKey() {
	spotify_key, spotify_app = config.GetSpotifyKeyAndApp()
	var err any
	token.AccessToken, err = getAccessToken(spotify_key, spotify_app)
	if err != nil {
		utils.LogError("Erreur de la récupération du token", err)
		return
	}

}

type TokenResponse struct {
	AccessToken string
}

func getAccessToken(clientID, clientSecret string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tokenRes TokenResponse
	json.NewDecoder(resp.Body).Decode(&tokenRes)
	return tokenRes.AccessToken, nil
}
