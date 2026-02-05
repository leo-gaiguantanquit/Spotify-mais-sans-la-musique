package route

import "html/template"

type HomeData struct {
	TopTitre_list template.HTML

	PP_mom_url              string
	TopArtist_name          string
	TopArtist_desc          string
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
	PromArtistURL           string
	PromArtist_track01_name string
	PromArtist_track01_time string
	PromArtist_track02_name string
	PromArtist_track02_time string
	PromArtist_track03_name string
	PromArtist_track03_time string
}
