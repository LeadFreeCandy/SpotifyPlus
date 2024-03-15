package main

import (
	"fmt"
	"time"

	"github.com/SpotifyPlus/internal"
	"github.com/SpotifyPlus/internal/scope"
	"go.uber.org/zap"
)

func main() {
	/*file, _ := os.OpenFile("internal/main/log.log", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	fileCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(file), zap.DebugLevel)*/

	config, _ := internal.NewConfigFromYaml("./config.yaml")
	logger, _ := zap.NewDevelopment()
	app := internal.NewApp(config, logger)

	app.InitializeAuthenticationRoute(nil)
	go func() {
		err := app.EnableHttpListener()
		if err != nil {
			logger.Error("Unexpected HTTP Listener Error", zap.Error(err))
		}
	}()

	url, _ := app.GenerateAuthenticationURL([]scope.Scope{scope.AppRemoteControl, scope.UgcImageUpload})
	fmt.Println(url.String())
	time.Sleep(10 * time.Second)

	artist, err := internal.GetArtist(&app, "5INjqkS1o8h1imAzPqGZBb")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(artist.Name)
	}
}
