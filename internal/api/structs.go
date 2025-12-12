package api

type SpotifyPlaylist struct {
	Items []struct {
		Track struct {
			Artists []struct {
				Name string
			}
			Name string
		}
	}
}
