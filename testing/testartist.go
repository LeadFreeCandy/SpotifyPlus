package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/SpotifyPlus/internal"
	"github.com/SpotifyPlus/internal/scope"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ServeURLToUser(url *url.URL) {
	fmt.Println(url.String())
}

func main() {
	// Prepare the app
	//loggerZap, _ := zap.NewDevelopment()
	file, _ := os.OpenFile("internal/main/log.log", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	fileCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(file), zap.DebugLevel)

	config, _ := internal.NewConfigFromYaml("internal/main/config.yaml")
	logger := zap.New(fileCore)
	app := internal.NewApp(config, logger)

	// Initialize auth route and listener
	app.InitializeAuthenticationRoute(nil)
	go func() {
		err := app.EnableHttpListener()
		if err != nil {
			logger.Error("Unexpected HTTP Listener Error", zap.Error(err))
		}
	}()

	// Prompt user to authenticate
	url, _ := app.GenerateAuthenticationURL([]scope.Scope{scope.AppRemoteControl, scope.UgcImageUpload})
	ServeURLToUser(url)
	artists, _ := internal.GetArtist(&app, "5INjqkS1o8h1imAzPqGZBb")
	fmt.Println(artists.Name)
}
