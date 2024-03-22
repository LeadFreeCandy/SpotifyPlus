package main

import (
	"fmt"
	"os"

	"github.com/SpotifyPlus/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	// create log file for zapcore
	file, _ := os.OpenFile("internal/main/test.log", os.O_CREATE|os.O_WRONLY, 0644)
	file.Close()
	file_core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(file), zap.DebugLevel)

	// get http config settings from Yaml
	yaml_config, _ := internal.NewConfigFromYaml("internal/main/config.yaml")
	log := zap.New(file_core)

	// initialize app, auth route, error listener
	app := internal.NewApp(yaml_config, log)
	app.InitializeAuthenticationRoute(file_core)
	err := app.EnableHttpListener()
	if err != nil {
		log.Error("http listener error: ", zap.Error(err))
	}

	// ask user for authentication, redirect to http handler
	// working on http listener for event handling
	auth_url, _ := app.GenerateAuthenticationURL([]scope.Scope{scope.AppRemoteControl, scope.UgcImageUpload, scope.UserReadRecentlyPlayed})
	fmt.Println("authenticate: ", auth_url)

	res -> auth_url
	/*
	var recent interface{}
	err := json.NewDecoder(res.Body).Decode(&recentlyPlayed)
	if err != nil {
		return nil, err
	}
	fmt.Println(recent)
	*/
}