package router

import (
	"demo-smtp/internal/api/controller"

	"github.com/gofiber/fiber/v2"
)

func CreateFiberInstance() *fiber.App {
	return fiber.New()
}

func ListenAndServe(app *fiber.App) error {
	emailController := controller.NewEmailController()
	templateController := controller.NewTemplateController()

	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api.Post("/send", emailController.SendPlainTextEmail)
	api.Post("/send/:slug", emailController.SendTemplateEmail)

	api.Get("/templates/:slug/raw", templateController.GetRaw)
	api.Post("/templates", templateController.Create)

	return app.Listen("localhost:8082")
}
