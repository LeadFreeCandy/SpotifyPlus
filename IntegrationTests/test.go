package main

import (
	"fmt"
	"os"
	"test/scope"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	file, _ := os.OpenFile("internal/main/log.log", os.O_CREATE|os.O_WRONLY, 0644)
	file.Close()
	file_core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(file), zap.DebugLevel)

	config, _ := internal.NewConfigFromYaml("internal/main/config.yaml")
	logger := zap.New(file_core)
	app := internal.NewApp(config, logger)

	app.InitializeAuthenticationRoute(nil)
	err := app.EnableHttpListener()
	if err != nil {
		logger.Error("Unexpected HTTP Listener Error", zap.Error(err))
	}

	url, _ := app.GenerateAuthenticationURL([]scope.Scope{scope.AppRemoteControl, scope.UserReadRecentlyPlayed})
	fmt.Println(url.String())
	time.Sleep(time.Hour)
}
