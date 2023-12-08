package main

import (
	"demo-smtp/internal/bootstrap"
)

// @title Demo SMTP API
// @description This is a demo SMTP API server.
// @version 1
// @host localhost:8082
// @BasePath /api

func main() {
	bootstrap.Start()
}
