package api

import (
	"SMSM/pkg/utils"
)

func Launch() {
	utils.Log("Chargement de l'API...")

	InitSpotify()
	// TODO: Ajouter les handlers nécéssaire pour js

	utils.Log("API chargés avec succé")
}
