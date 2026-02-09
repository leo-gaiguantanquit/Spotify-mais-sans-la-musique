package config

import (
	"SMSM/pkg/utils"
	"encoding/json"
	"os"
)

var cfg Config

// Config stocke la configuration globale de l'application.
type Config struct {
	API   ConfigAPI
	DEBUG bool
}

// ConfigAPI stocke les clés API pour Spotify et Ticketmaster.
type ConfigAPI struct {
	Spotify_key      string
	Spotify_app      string
	Ticketmaster_key string
}

// InitConfig charge la configuration depuis le fichier config.json
// et définit le mode debug si nécessaire.
func InitConfig() {
	utils.Log("Initialisation des configs...")

	file, err := os.Open("config.json")
	if err != nil {
		utils.LogError("Erreur de l'initialisation config", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		utils.LogError("Erreur lors du décodage JSON", err)
	}

	utils.Log("Configs initialisés avec succée")
	utils.SetDebugMode(GetDebug())
}

// GetSpotifyKeyAndApp retourne le Client Secret et le Client ID de l'API Spotify.
func GetSpotifyKeyAndApp() (string, string) {
	return cfg.API.Spotify_key, cfg.API.Spotify_app
}

// GetTicketmasterKey retourne la clé API de Ticketmaster.
func GetTicketmasterKey() string {
	return cfg.API.Ticketmaster_key
}

// GetDebug retourne l'état du mode debug (vrai ou faux).
func GetDebug() bool {
	return cfg.DEBUG
}
