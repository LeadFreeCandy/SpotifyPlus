package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type Artist struct {
	ExternalUrls struct {
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

func GetArtist(app *AppState, id string) (*Artist, error) {
	//generate request url
	reqURL := (&url.URL{}).JoinPath("/artists").JoinPath(id)

	//generate API request, serve to http client
	req, err := app.generateApiRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		app.logger.Error("Failed to generate request", zap.Error(err))
		return nil, err
	}

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		app.logger.Error("Client failed request", zap.Error(err))
		return nil, err
	}
	println(res.Status)

	//decode json response, return artist
	decoder := json.NewDecoder(res.Body)
	artist := &Artist{}
	err = decoder.Decode(artist)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return artist, nil
}
