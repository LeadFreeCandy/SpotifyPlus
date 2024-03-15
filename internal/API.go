package internal

import (
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type SimplifiedImage struct {
	URL    string `json:"url"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type SimplifiedArtist struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href string `json:"href"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type SimplifiedTrack struct {
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []string           `json:"available_markets"`
	DiscNumber       int                `json:"disc_number"`
	DurationMs       int                `json:"duration_ms"`
	Explicit         bool               `json:"explicit"`

	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`

	Href       string `json:"href"`
	ID         string `json:"id"`
	IsPlayable bool   `json:"is_playable"`

	LinkedFrom struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Id   string `json:"id"`
		Type string `json:"type"`
		Uri  string `json:"uri"`
	} `json:"linked_from"`

	Restrictions struct {
		Reason string `json:"reason"`
	} `json:"restrictions"`
	Name        string `json:"name"`
	PreviewURL  string `json:"preview_url"`
	TrackNumber int    `json:"track_number"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	IsLocal     bool   `json:"is_local"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

func (app *AppState) generateApiRequest(method string, path *url.URL, body io.Reader) (*http.Request, error) {
	baseAPI, _ := url.Parse("https://api.spotify.com/v1")

	app.logger.Info("Base api", zap.String("baseAPI", baseAPI.String()))
	apiURI := baseAPI.JoinPath(path.Path)

	apiURI.RawQuery = path.RawQuery
	app.logger.Info("api request", zap.String("url", apiURI.String()))

	request, err := http.NewRequest(method, apiURI.String(), body)

	if err != nil {
		return nil, err
	}

	headers := request.Header
	headers.Add("Authorization", "Bearer "+app.authenticationToken)

	request.Header = headers
	return request, nil
}
