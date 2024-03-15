package internal

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Album struct {
	AlbumType        string   `json:"album_type"`
	TotalTracks      int      `json:"total_tracks"`
	AvailableMarkets []string `json:"available_markets"`

	ExternalUrls struct { // Non-named objects from the spotify api are going to be raw structs for now
		Spotify string `json:"spotify"`
	} `json:"external_urls"`

	Href                 string            `json:"href"`
	ID                   string            `json:"id"`
	Images               []SimplifiedImage `json:"images"`
	Name                 string            `json:"name"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`

	Restrictions struct {
		Reason string `json:"reason"`
	} `json:"restrictions"`

	Type    string             `json:"type"`
	URI     string             `json:"uri"`
	Artists []SimplifiedArtist `json:"artists"`

	Tracks struct {
		Href     string            `json:"href"`
		Limit    int               `json:"limit"`
		Next     string            `json:"next"`
		Offset   int               `json:"offset"`
		Previous string            `json:"previous"`
		Total    int               `json:"total"`
		Items    []SimplifiedTrack `json:"items"`
	} `json:"tracks"`

	Copyrights []Copyright `json:"copyrights"`

	ExternalIds struct {
		Isrc string `json:"isrc"`
		Ean  string `json:"ean"`
		Upc  string `json:"upc"`
	} `json:"external_ids"`

	Genres     []string `json:"genres"`
	Label      string   `json:"label"`
	Popularity int      `json:"popularity"`
}

func GetAlbum(app *AppState, id string, market string) (*Album, error) {
	requestURL := (&url.URL{}).JoinPath("/albums").JoinPath(id)
	// requestURL = "/albums/1233123

	if market != "" {
		query := requestURL.Query()
		query.Add("market", market)
		requestURL.RawQuery = query.Encode()
		// requestURL = "/albums/1233123?market=EN"
	}

	request, err := app.generateApiRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	result, err := (&http.Client{}).Do(request)
	if err != nil {
		return nil, err
	}

	defer result.Body.Close()

	decoder := json.NewDecoder(result.Body)
	album := &Album{}
	err = decoder.Decode(album)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func GetAlbums(app *AppState, ids []string, market string) ([]*Album, error) {
	requestPath := (&url.URL{}).JoinPath("/albums")
	commaSeperatedIds := strings.Join(ids, ",")
	query := requestPath.Query()
	query.Add("id", commaSeperatedIds)

	if market != "" {
		query.Add("market", market)
	}

	requestPath.RawQuery = query.Encode()

	request, err := app.generateApiRequest(http.MethodGet, requestPath, nil)
	if err != nil {
		return nil, err
	}

	result, err := (&http.Client{}).Do(request)
	if err != nil {
		return nil, err
	}

	defer result.Body.Close()

	jsonDecoder := json.NewDecoder(result.Body)
	albumsList := struct {
		Albums []*Album `json:"albums"`
	}{}
	err = jsonDecoder.Decode(&albumsList)
	if err != nil {
		return nil, err
	}

	return albumsList.Albums, nil
}
