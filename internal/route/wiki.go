package route

import (
	"SMSM/internal/api"
	"SMSM/pkg/utils"
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	TabTitle string
	Body     template.HTML
	Header   template.HTML
	Footer   template.HTML
}

func home(w http.ResponseWriter, r *http.Request) {
	topArtistID := "5y239Fviw0hAOe5Jy3qzlf"

	utils.Log("Page Home ouverte")

	playlist, err := api.GetTopArtist()
	if err != nil {
		utils.LogError("Erreur api spotify", err)
	}
	topArtiste, err := api.GetArtist(topArtistID)
	if err != nil {
		utils.LogError("Erreur api spotify", err)
	}
	topArtisteTrack, err := api.GetArtistTopTrack(topArtistID)
	// utils.Debug(fmt.Sprintf("%v", playlist))
	// utils.Debug(fmt.Sprintf("%v", topArtiste))
	tracksHTML := ""

	if err != nil {
		utils.LogError("Erreur api spotify", err)
	} else {
		for _, item := range playlist.Tracks.Items {
			imgUrl := "https://placehold.co/200x200"
			if len(item.Track.Album.Images) > 0 {
				imgUrl = item.Track.Album.Images[0].Url
			}

			artistName := ""
			if len(item.Track.Artists) > 0 {
				artistName = item.Track.Artists[0].Name
			}

			tracksHTML += render("layer/trans-pp-topTitre", map[string]interface{}{
				"PP_url":     imgUrl,
				"Name":       item.Track.Name,
				"ArtistName": artistName,
			})
		}
	}

	data := HomeData{
		TopTitre_list: template.HTML(tracksHTML),

		// Données statiques pour la démo "Artiste du Moment"
		PP_mom_url:             topArtiste.Images[0].URL,
		TopArtist_name:         topArtiste.Name,
		TopArtist_desc:         strings.Join(topArtiste.Genres, " - "),
		TopArtist_track01_name: topArtisteTrack.Tracks[0].Name,
		TopArtist_track01_time: utils.MsToTime(topArtisteTrack.Tracks[0].Duration),
		TopArtist_track02_name: topArtisteTrack.Tracks[1].Name,
		TopArtist_track02_time: utils.MsToTime(topArtisteTrack.Tracks[1].Duration),
		TopArtist_track03_name: topArtisteTrack.Tracks[2].Name,
		TopArtist_track03_time: utils.MsToTime(topArtisteTrack.Tracks[2].Duration),
		PP_prom_url:            "https://i.scdn.co/image/ab67616d0000b273bf75176711956555132dd0e2",
	}

	renderFile("Accueil", render("home", data), w)
}

func renderFile(title string, content string, w http.ResponseWriter) {
	header := render("header", nil)
	footer := render("footer", nil)
	data := PageData{
		TabTitle: title,
		Body:     template.HTML(content),
		Header:   template.HTML(header),
		Footer:   template.HTML(footer),
	}

	executeHTML("layout", data, w)
}

func executeHTML(fileName string, data any, w http.ResponseWriter) {
	pathString := fmt.Sprintf("template/%v.html", fileName)

	t, _ := template.ParseFiles(pathString)
	t.Execute(w, data)
}

func render(fileName string, data any) string {
	pathString := fmt.Sprintf("template/%v.html", fileName)

	t, _ := template.ParseFiles(pathString)
	var buf bytes.Buffer
	t.Execute(&buf, data)
	return buf.String()
}
