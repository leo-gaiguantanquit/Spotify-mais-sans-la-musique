package route

import "html/template"

// HomeData contient les données affichées sur la page d'accueil,
// notamment les artistes populaires et les suggestions.
type HomeData struct {
	TopTitre_list template.HTML

	PP_mom_url              string
	TopArtist_name          string
	TopArtist_desc          string
	TopArtist_ID            string
	TopArtist_track01_name  string
	TopArtist_track01_time  string
	TopArtist_track02_name  string
	TopArtist_track02_time  string
	TopArtist_track03_name  string
	TopArtist_track03_time  string
	TopArtistURL            string
	PP_prom_url             string
	PromArtist_name         string
	PromArtist_desc         string
	PromArtist_ID           string
	PromArtistURL           string
	PromArtist_track01_name string
	PromArtist_track01_time string
	PromArtist_track02_name string
	PromArtist_track02_time string
	PromArtist_track03_name string
	PromArtist_track03_time string
}

// TrackData structure les informations détaillées d'une piste musicale.
type TrackData struct {
	Rank     int
	Title    string
	Duration string
}

// ConcertData structure les informations détaillées d'un concert.
type ConcertData struct {
	Name      string
	Date      string
	City      string
	Venue     string
	TicketUrl string
	Lat       string
	Lng       string
}

// ArtistData contient toutes les informations nécessaires à l'affichage de la page d'un artiste.
type ArtistData struct {
	Artist_img_url    string
	Artist_name       string
	Artist_desc       string
	Artist_bio        template.HTML
	Artist_top_tracks []TrackData
	Artist_concerts   []ConcertData
	Artist_url        string
}

// ArtistListItem représente un artiste dans une liste simplifiée (ex: résultats de recherche).
type ArtistListItem struct {
	Image string
	Nom   string
	Genre string
	ID    string
}

// ArtistListData contient une liste d'artistes à afficher.
type ArtistListData struct {
	Artistes []ArtistListItem
}

// ConcertPageData contient la liste des concerts à afficher sur la page concerts.
type ConcertPageData struct {
	Events []ConcertData
}
