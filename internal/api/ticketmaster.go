package api

import (
	"SMSM/internal/config"
	"SMSM/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type TicketmasterEventResponse struct {
	Embedded struct {
		Events []struct {
			Name  string `json:"name"`
			Dates struct {
				Start struct {
					LocalDate string `json:"localDate"`
				} `json:"start"`
			} `json:"dates"`
			Embedded struct {
				Venues []struct {
					Name string `json:"name"`
					City struct {
						Name string `json:"name"`
					} `json:"city"`
					Country struct {
						Name        string `json:"name"`
						CountryCode string `json:"countryCode"`
					} `json:"country"`
					Location struct {
						Longitude string `json:"longitude"`
						Latitude  string `json:"latitude"`
					} `json:"location"`
				} `json:"venues"`
			} `json:"_embedded"`
			Url string `json:"url"`
		} `json:"events"`
	} `json:"_embedded"`
}

func GetArtistEvents(artistName string) (*TicketmasterEventResponse, error) {
	apiKey := config.GetTicketmasterKey()
	if apiKey == "" {
		return nil, fmt.Errorf("ticketmaster api key is missing")
	}

	utils.Debug(fmt.Sprintf("Fetching events for artist: %s", artistName))

	// Ticketmaster API endpoint
	baseURL := "https://app.ticketmaster.com/discovery/v2/events.json"
	params := url.Values{}
	params.Add("apikey", apiKey)
	params.Add("keyword", artistName)
	params.Add("sort", "date,asc")
	params.Add("size", "15")

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		utils.LogError("Error making request to Ticketmaster", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ticketmaster api error: %d", resp.StatusCode)
	}

	var result TicketmasterEventResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		utils.LogError("Error decoding Ticketmaster response", err)
		return nil, err
	}

	return &result, nil
}

func GetGenericEvents() (*TicketmasterEventResponse, error) {
	apiKey := config.GetTicketmasterKey()
	if apiKey == "" {
		return nil, fmt.Errorf("ticketmaster api key is missing")
	}

	utils.Debug("Fetching generic music events in France")

	// Ticketmaster API endpoint
	baseURL := "https://app.ticketmaster.com/discovery/v2/events.json"
	params := url.Values{}
	params.Add("apikey", apiKey)
	params.Add("classificationName", "Music")
	//params.Add("countryCode", "FR")
	params.Add("latlong", "48.8566,2.3522") // Paris
	params.Add("radius", "1000")
	params.Add("unit", "km")
	params.Add("sort", "date,asc")
	params.Add("size", "50")

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		utils.LogError("Error making request to Ticketmaster", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ticketmaster api error: %d", resp.StatusCode)
	}

	var result TicketmasterEventResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		utils.LogError("Error decoding Ticketmaster response", err)
		return nil, err
	}

	return &result, nil
}
