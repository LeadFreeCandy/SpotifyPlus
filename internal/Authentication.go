package internal

import (
	"fmt"
	"github.com/SpotifyPlus/internal/scope"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"net/url"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// InitializeAuthenticationRoute Adds authentication handling to the server router.
// It exposes an extraHandler parameter which is optional, which will be executed at the end.
// If provided the extraHandler is expected to handle generating a valid httpResponse
func (app *AppState) InitializeAuthenticationRoute(extraHandler func(http.ResponseWriter, *http.Request)) error {
	redirectURI, err := url.Parse(app.config.RedirectURI)
	extraHandlerExists := extraHandler != nil
	app.logger.Info("Initializing Authentication Route", zap.Bool("ExtraHandlerExists", extraHandlerExists))
	if err != nil {
		app.logger.Error("URI", zap.Error(err))
		return err
	}
	app.serverMux.HandleFunc(redirectURI.Path, func(writer http.ResponseWriter, request *http.Request) {
		logger := app.logger.With(
			zap.String("service", "Auth Callback"),
			zap.String("request", request.RequestURI),
		)
		logger.Info("Received authentication request")

		err := app.processAuthenticationRequestParams(request)
		if err != nil {
			logger.Error("Failed to process request", zap.Error(err))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		// If the website team wants to use this function they can control where to redirect after successful auth
		if extraHandler != nil {
			logger.Info("Passing to extraHandler")
			extraHandler(writer, request)
			return
		}
		// Else we just return a simple OK
		writer.WriteHeader(http.StatusOK)
	})

	return nil
}

func (app *AppState) processAuthenticationRequestParams(request *http.Request) error {
	queryParams := request.URL.Query()
	state := queryParams.Get("state")
	code := queryParams.Get("code")
	if state != app.authenticationState {
		return fmt.Errorf("Invalid state. Expected: %v \t Actual: %v", app.authenticationState, state)
	}
	app.authenticationToken = code

	return nil
}

func (app *AppState) GenerateAuthenticationURL(scopes []scope.Scope) (*url.URL, error) {
	state := generateRandomString(16)

	responseURL, err := url.Parse("https://accounts.spotify.com/authorize")

	if err != nil { // Should never happen
		app.logger.Error("Failed to parse authorize url", zap.Error(err))
		return responseURL, err
	}

	// Update app state for CSRF
	app.authenticationState = state

	queryParams := responseURL.Query()
	queryParams.Add("client_id", app.config.ClientID)
	queryParams.Add("response_type", "code")
	queryParams.Add("redirect_uri", app.config.RedirectURI)
	queryParams.Add("state", state) // Optional

	for _, providedScope := range scopes {
		queryParams.Add("scope", string(providedScope))
	}

	responseURL.RawQuery = queryParams.Encode()
	return responseURL, nil
}
