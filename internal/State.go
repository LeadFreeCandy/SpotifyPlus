package internal

import (
	"net/http"
	"strconv"
)

type AppState struct {
	config              Config
	server              http.Server
	serverMux           *http.ServeMux
	authenticationState string
	authenticationToken string
}

func NewApp(config Config) AppState {
	mux := http.NewServeMux()

	return AppState{
		config: config,
		server: http.Server{
			Addr: ":" + strconv.FormatInt(int64(config.serverPort), 10),
		},
		serverMux: mux,
	}
}

func (app *AppState) EnableHttpListener() error {
	app.server.Handler = app.serverMux
	err := app.server.ListenAndServe()
	return err
}

func (app *AppState) CloseHttpListener() error {
	err := app.server.Close()
	return err
}
