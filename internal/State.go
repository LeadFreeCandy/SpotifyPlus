package internal

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type AppState struct {
	config              Config
	server              http.Server
	serverMux           *http.ServeMux
	authenticationState string
	authenticationToken string
	logger              *zap.Logger
}

func NewApp(config Config, logger *zap.Logger) AppState {
	mux := http.NewServeMux()

	return AppState{
		config: config,
		server: http.Server{
			Addr: ":" + strconv.FormatInt(int64(config.ServerPort), 10),
		},
		serverMux: mux,
		logger:    logger,
	}
}

func (app *AppState) EnableHttpListener() error {
	app.server.Handler = app.serverMux
	app.logger.Info("Listening for http connections")
	err := app.server.ListenAndServe()
	return err
}

func (app *AppState) CloseHttpListener() error {
	app.logger.Info("Stopped listening for http connections")
	err := app.server.Close()
	return err
}
