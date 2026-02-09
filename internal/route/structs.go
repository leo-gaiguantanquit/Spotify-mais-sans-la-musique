package route

import "html/template"

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

type TrackData struct {
	Rank     int
	Title    string
	Duration string
}

type ConcertData struct {
	Name      string
	Date      string
	City      string
	Venue     string
	TicketUrl string
	Lat       string
	Lng       string
}

type ArtistData struct {
	Artist_img_url    string
	Artist_name       string
	Artist_desc       string
	Artist_bio        template.HTML
	Artist_top_tracks []TrackData
	Artist_concerts   []ConcertData
	Artist_url        string
}

type ArtistListItem struct {
	Image string
	Nom   string
	Genre string
	ID    string
}

type ArtistListData struct {
	Artistes []ArtistListItem
}

type ConcertPageData struct {
	Events []ConcertData
}
