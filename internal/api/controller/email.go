package controller

import (
	"demo-smtp/internal/api"
	"demo-smtp/internal/config"
	newSmtp "demo-smtp/internal/smtp"
	"demo-smtp/internal/template"
	"demo-smtp/internal/types"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type EmailControllerInterface interface {
	SendPlainTextEmail(c *fiber.Ctx) error
	SendTemplateEmail(c *fiber.Ctx) error
}

type EmailController struct {
	EmailControllerInterface
	validate        *validator.Validate
	templateService template.TemplateServiceInterface
}

func NewEmailController() *EmailController {
	return &EmailController{
		validate:        validator.New(),
		templateService: template.NewTemplateService(),
	}
}

// Handler functions
// SendPlainTextEmail godoc
// @Summary Send plain text email
// @Description Send plain text email
// @Tags Email
// @Accept json
// @Produce json
// @Param body body types.PlainTextEmail true "body"
// @Success 201 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /send [post]
func (e *EmailController) SendPlainTextEmail(c *fiber.Ctx) error {
	var bodyEmail types.PlainTextEmail

	if err := c.BodyParser(&bodyEmail); err != nil {
		return api.Err(c, fiber.StatusBadRequest, "Invalid body 1", err)
	}

	if err := e.validate.Struct(bodyEmail); err != nil {
		return api.Err(c, fiber.StatusBadRequest, "Invalid body 2", err)
	}

	mail := types.Mail{
		To:      []string{bodyEmail.To},
		Subject: bodyEmail.Subject,
		Sender:  config.MainConfig.DefaultFrom,
		Body:    bodyEmail.Body,
		Type:    types.TEXT,
	}

	// newSmtp.SendEmail(&mail)
	err := newSmtp.SendEmail(&mail)

	if err != nil {
		fmt.Println("Error sending email", err)
		return api.Err(c, fiber.StatusInternalServerError, "Error getting queue", err)
	}

	api.Success(c, fiber.StatusCreated, "Email sent successfully")
	return nil
}

// SendTemplateEmail godoc
// @Summary Send template email
// @Description Send template email
// @Tags Email
// @Accept json
// @Produce json
// @Param slug path string true "Template slug"
// @Param body body types.TemplateEmail true "body"
// @Success 201 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /send/{slug} [post]
func (e *EmailController) SendTemplateEmail(ctx *fiber.Ctx) error {
	var templateEmail types.TemplateEmail

	templateName := ctx.Params("slug")

	if templateName == "" {
		return api.Err(ctx, fiber.StatusBadRequest, "Invalid template name", nil)
	}

	if err := ctx.BodyParser(&templateEmail); err != nil {
		return api.Err(ctx, fiber.StatusBadRequest, "Invalid body", err)
	}

	if err := e.validate.Struct(templateEmail); err != nil {
		return api.Err(ctx, fiber.StatusBadRequest, "Invalid body", err)
	}

	template, err := e.templateService.ParseTemplate(templateName, templateEmail.Data)

	if err != nil {
		return api.Err(ctx, fiber.StatusInternalServerError, "Error parsing template", err)
	}

	mail := types.Mail{
		To:      []string{templateEmail.To},
		Subject: templateEmail.Subject,
		Sender:  config.MainConfig.DefaultFrom,
		Body:    template.String(),
		Type:    types.HTML,
	}

	err = newSmtp.SendEmail(&mail)

	if err != nil {
		return api.Err(ctx, fiber.StatusInternalServerError, "Error getting queue", err)
	}

	api.Success(ctx, fiber.StatusCreated, "Email sent successfully")
	return nil
}
