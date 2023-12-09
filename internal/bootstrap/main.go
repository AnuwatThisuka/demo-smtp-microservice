package bootstrap

import (
	"demo-smtp/internal/api/router"
	"demo-smtp/internal/config"
	newSmtp "demo-smtp/internal/smtp"
	"fmt"
	"log/slog"

	"github.com/streadway/amqp"
)

func Start() {

	err := config.LoadConfig()

	if err != nil {
		slog.Error("Failed to load envs: " + err.Error())
	}

	fmt.Println("RabbitMQ in Golang: Getting started tutorial")

	connection, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%d/",
			config.MainConfig.RabbitMQ.Username,
			config.MainConfig.RabbitMQ.Password,
			config.MainConfig.RabbitMQ.Host,
			config.MainConfig.RabbitMQ.Port,
		),
	)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ instance")

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
