package main

import (
	"SMSM/internal/api"
	"SMSM/internal/config"
	"SMSM/internal/route"
)

// main est le point d'entrée de l'application.
// Il initialise la configuration, lance les services API et démarre le serveur HTTP.
func main() {
	config.InitConfig()
	api.Launch()
	route.Launch()
}
