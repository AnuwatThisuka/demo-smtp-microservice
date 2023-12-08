package bootstrap

import (
	"demo-smtp/internal/api/router"
	"demo-smtp/internal/config"
	newSmtp "demo-smtp/internal/smtp"
	"log/slog"
)

func Start() {

	err := config.LoadConfig()

	if err != nil {
		slog.Error("Failed to load envs: " + err.Error())
	}

	SetupLog()
	err = newSmtp.Ping()
	if err != nil {
		slog.Error(err.Error())
	}

	app := router.CreateFiberInstance()

	err = router.ListenAndServe(app)

	if err != nil {
		slog.Error("Failed to start HTTP server: " + err.Error())
	}

}
