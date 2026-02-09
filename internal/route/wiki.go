package route

import (
	"SMSM/internal/api"
	"SMSM/pkg/utils"
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type PageData struct {
	TabTitle string
	Body     template.HTML
	Header   template.HTML
	Footer   template.HTML
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderFile("Cieelll - 404", render("404", nil), w)
		return
	}

	topArtistID := "5y239Fviw0hAOe5Jy3qzlf"
	promArtistID := "2pGLZYyPlXH1rkOsxRdZ0h"

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
	if err != nil {
		utils.LogError("Erreur api spotify", err)
	}

	promArtiste, err := api.GetArtist(promArtistID)
	if err != nil {
		utils.LogError("Erreur api spotify", err)
	}
	promArtisteTrack, err := api.GetArtistTopTrack(promArtistID)
	if err != nil {
		utils.LogError("Erreur api spotify", err)
	}

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
		TopArtistURL:           topArtiste.ExternalURLs.Spotify,

		PP_prom_url:             promArtiste.Images[0].URL,
		PromArtist_name:         promArtiste.Name,
		PromArtist_desc:         strings.Join(promArtiste.Genres, " - "),
		PromArtist_track01_name: promArtisteTrack.Tracks[0].Name,
		PromArtist_track01_time: utils.MsToTime(promArtisteTrack.Tracks[0].Duration),
		PromArtist_track02_name: promArtisteTrack.Tracks[1].Name,
		PromArtist_track02_time: utils.MsToTime(promArtisteTrack.Tracks[1].Duration),
		PromArtist_track03_name: promArtisteTrack.Tracks[2].Name,
		PromArtist_track03_time: utils.MsToTime(promArtisteTrack.Tracks[2].Duration),
		PromArtistURL:           promArtiste.ExternalURLs.Spotify,
	}

	renderFile("Accueil", render("home", data), w)
}

func artist(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		renderFile("Cieelll - 404", render("404", nil), w)
		return
	}
	artistID := parts[2]

	artist, err := api.GetArtist(artistID)
	if err != nil {
		renderFile("Cieelll - 404", render("404", nil), w)
		return
	}

	topTracks, err := api.GetArtistTopTrack(artistID)
	if err != nil {
		utils.LogError("Error getting top tracks", err)
	}

	var tracks []TrackData
	if topTracks != nil {
		for i, track := range topTracks.Tracks {
			if i >= 5 {
				break
			} // Top 5
			tracks = append(tracks, TrackData{
				Rank:     i + 1,
				Title:    track.Name,
				Duration: utils.MsToTime(track.Duration),
			})
		}
	}

	imageUrl := "https://placehold.co/600x400"
	if len(artist.Images) > 0 {
		imageUrl = artist.Images[0].URL
	}

	description := "Genre non spécifié"
	if len(artist.Genres) > 0 {
		description = strings.Join(artist.Genres, " - ")
	}

	bioPlaceholder := fmt.Sprintf("Retrouvez %s sur Spotify pour plus d'informations. %d abonnés suivent cet artiste.", artist.Name, artist.Followers.Total)

	// Fetch Concerts
	events, err := api.GetArtistEvents(artist.Name)
	var concerts []ConcertData
	if err != nil {
		utils.LogError("Error getting artist events", err)
	} else if events != nil && len(events.Embedded.Events) > 0 {
		for _, event := range events.Embedded.Events {
			venueName := ""
			cityName := ""
			lat := ""
			lng := ""

			if len(event.Embedded.Venues) > 0 {
				venueName = event.Embedded.Venues[0].Name
				cityName = event.Embedded.Venues[0].City.Name + ", " + event.Embedded.Venues[0].Country.Name
				lat = event.Embedded.Venues[0].Location.Latitude
				lng = event.Embedded.Venues[0].Location.Longitude
			}

			// Format date
			dateStr := event.Dates.Start.LocalDate
			parsedDate, _ := time.Parse("2006-01-02", dateStr)
			formattedDate := parsedDate.Format("02 Jan") // "12 Oct"

			concerts = append(concerts, ConcertData{
				Name:      event.Name,
				Date:      formattedDate,
				City:      cityName,
				Venue:     venueName,
				TicketUrl: event.Url,
				Lat:       lat,
				Lng:       lng,
			})
		}
	}

	data := ArtistData{
		Artist_img_url:    imageUrl,
		Artist_name:       artist.Name,
		Artist_desc:       description,
		Artist_bio:        template.HTML(bioPlaceholder),
		Artist_top_tracks: tracks,
		Artist_concerts:   concerts,
	}

	renderFile(artist.Name, render("artiste", data), w)
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

	t, err := template.ParseFiles(pathString)
	if err != nil {
		utils.LogError(fmt.Sprintf("Error parsing template %s", fileName), err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		utils.LogError(fmt.Sprintf("Error executing template %s", fileName), err)
	}
}

func render(fileName string, data any) string {
	pathString := fmt.Sprintf("template/%v.html", fileName)

	t, err := template.ParseFiles(pathString)
	if err != nil {
		utils.LogError(fmt.Sprintf("Error parsing template %s", fileName), err)
		return "Error rendering template"
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		utils.LogError(fmt.Sprintf("Error executing template %s", fileName), err)
		return "Error executing template"
	}
	return buf.String()
}

func allArtistes(w http.ResponseWriter, r *http.Request) {
	utils.Log("Page Artistes ouverte")

	genre := r.URL.Query().Get("genre")
	searchQuery := "genre:pop"
	if genre != "" {
		searchQuery = "genre:" + genre
	}

	searchRes, err := api.Search(searchQuery)
	if err != nil {
		utils.LogError("Erreur recherche artistes", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}

	var artistesList []ArtistListItem
	if searchRes != nil {
		for _, item := range searchRes.Artists.Items {
			imgUrl := "https://placehold.co/200x200"
			if len(item.Images) > 0 {
				imgUrl = item.Images[0].URL
			}

			genres := "Inconnu"
			if len(item.Genres) > 0 {
				genres = strings.Join(item.Genres, ", ")
			}

			artistesList = append(artistesList, ArtistListItem{
				Image: imgUrl,
				Nom:   item.Name,
				Genre: genres,
				ID:    item.ID,
			})
		}
	}

	data := ArtistListData{
		Artistes: artistesList,
	}

	renderFile("Artistes", render("artistes-listes", data), w)
}

func concert(w http.ResponseWriter, r *http.Request) {
	utils.Log("Page Concerts ouverte")

	events, err := api.GetGenericEvents()

	var concerts []ConcertData
	if err != nil {
		utils.LogError("Error getting generic events", err)
	} else if events != nil && len(events.Embedded.Events) > 0 {
		for _, event := range events.Embedded.Events {
			venueName := ""
			cityName := ""
			lat := ""
			lng := ""

			if len(event.Embedded.Venues) > 0 {
				if event.Embedded.Venues[0].Country.CountryCode != "FR" {
					continue
				}

				venueName = event.Embedded.Venues[0].Name
				cityName = event.Embedded.Venues[0].City.Name
				if event.Embedded.Venues[0].Country.Name != "" {
					cityName += ", " + event.Embedded.Venues[0].Country.Name
				}
				lat = event.Embedded.Venues[0].Location.Latitude
				lng = event.Embedded.Venues[0].Location.Longitude
			} else {
				continue
			}

			// Filter out invalid coordinates (Ticketmaster returns "0","0")
			if lat == "0" || lng == "0" {
				lat = ""
				lng = ""
			}

			// Format date
			dateStr := event.Dates.Start.LocalDate
			parsedDate, _ := time.Parse("2006-01-02", dateStr)
			formattedDate := parsedDate.Format("02 Jan")

			concerts = append(concerts, ConcertData{
				Name:      event.Name,
				Date:      formattedDate,
				City:      cityName,
				Venue:     venueName,
				TicketUrl: event.Url,
				Lat:       lat,
				Lng:       lng,
			})
		}
	}

	data := ConcertPageData{
		Events: concerts,
	}

	renderFile("Concert", render("concert", data), w)
}

func search(w http.ResponseWriter, r *http.Request) {
	utils.Log("Recherche")

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Redirect(w, r, "/artistes", http.StatusFound)
		return
	}

	searchRes, err := api.Search(query)
	if err != nil {
		utils.LogError("Erreur recherche artistes", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}

	var artistesList []ArtistListItem
	if searchRes != nil {
		for _, item := range searchRes.Artists.Items {
			imgUrl := "https://placehold.co/200x200"
			if len(item.Images) > 0 {
				imgUrl = item.Images[0].URL
			}

			genres := "Inconnu"
			if len(item.Genres) > 0 {
				genres = strings.Join(item.Genres, ", ")
			}

			artistesList = append(artistesList, ArtistListItem{
				Image: imgUrl,
				Nom:   item.Name,
				Genre: genres,
				ID:    item.ID,
			})
		}
	}

	data := ArtistListData{
		Artistes: artistesList,
	}

	renderFile("Recherche : "+query, render("artistes-listes", data), w)
}
