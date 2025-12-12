package main

import (
	"SMSM/internal/api"
	"SMSM/internal/route"
)

func main() {
	route.Launch()
	api.Launch()
}
