package api

// SpotifyPlaylist structure la réponse d'une playlist (liste d'items).
type SpotifyPlaylist struct {
	Name   string `json:"name"`
	Tracks struct {
		Items []struct {
			Track struct {
				Name    string `json:"name"`
				Artists []struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"artists"`
				Album struct {
					Images []struct {
						Url string `json:"url"`
					} `json:"images"`
				} `json:"album"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
}

// SpotifyArtist représente un artiste et ses détails.
type SpotifyArtist struct {
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"followers"`
	Genres []string `json:"genres"`
	Href   string   `json:"href"`
	ID     string   `json:"id"`
	Images []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
	Type       string `json:"type"`
	URI        string `json:"uri"`
}

// SpotifyArtistTopTracks contient la liste des titres les plus populaires d'un artiste.
type SpotifyArtistTopTracks struct {
	Tracks []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Duration int    `json:"duration_ms"`
		Album    struct {
			Name   string `json:"name"`
			Images []struct {
				URL string `json:"url"`
			} `json:"images"`
		} `json:"album"`
	} `json:"tracks"`
}

// SpotifySearchResponse structure la réponse d'une recherche (playlists, artistes, titres).
type SpotifySearchResponse struct {
	Playlists struct {
		Items []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"items"`
	} `json:"playlists"`
	Artists struct {
		Items []SpotifyArtist `json:"items"`
	} `json:"artists"`
	Tracks struct {
		Items []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			Album struct {
				Images []struct {
					URL string `json:"url"`
				} `json:"images"`
			} `json:"album"`
		} `json:"items"`
	} `json:"tracks"`
}
