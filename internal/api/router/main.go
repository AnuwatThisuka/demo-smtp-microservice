package router

import (
	"demo-smtp/internal/api/controller"

	_ "demo-smtp/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func CreateFiberInstance() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	return app
}

func ListenAndServe(app *fiber.App) error {
	emailController := controller.NewEmailController()
	templateController := controller.NewTemplateController()

	// swagger:operation POST /send Email SendPlainTextEmail

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	api.Post("/send", emailController.SendPlainTextEmail)
	api.Post("/send/:slug", emailController.SendTemplateEmail)

	api.Get("/templates/:slug/raw", templateController.GetRaw)
	api.Post("/templates", templateController.Create)

	return app.Listen("localhost:8082")
}
