package config

import (
	"SMSM/pkg/utils"
	"encoding/json"
	"os"
)

var (
	cfg Config
)

type Config struct {
	API ConfigAPI
}

type ConfigAPI struct {
	Spotify_key string
	Spotify_app string
}

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
}

func GetSpotifyKeyAndApp() (string, string) {
	return cfg.API.Spotify_key, cfg.API.Spotify_app
}
