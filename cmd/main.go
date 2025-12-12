package main

import (
	"SMSM/internal/api"
	"SMSM/internal/config"
	"SMSM/internal/route"
)

func main() {
	config.InitConfig()
	api.Launch()
	route.Launch()
}
