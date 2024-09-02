package controllers

import (
	"golang-boilerplate/platform/jwt"
	"golang-boilerplate/service/auth"
	"golang-boilerplate/service/reconcilliation"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	AuthService  auth.AuthServiceInterface
	ReconService reconcilliation.ReconciliationServiceInterface
	AuthConfig   jwt.Config
}

func (c Controller) Serve() {
	f := fiber.New()

	f.Get("/", healthCheck)

	authGroup := f.Group("/auth")

	authGroup.Post("/login", c.login)

	reconciliationGroup := f.Group("/reconciliation")
	reconciliationGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(c.AuthConfig.SigningKey),
		},
	}))
	reconciliationGroup.Post("/", c.reconciliation)

	f.Listen(":3000")
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "OK",
	})
}
