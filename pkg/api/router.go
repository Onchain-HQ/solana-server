package api

import (
	handler "github.com/Onchain-HQ/solana-server/pkg/handlers"
	"github.com/gofiber/fiber/v2"
)

func New(router fiber.Router, handler *handler.Handler) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	router.Post("/address", handler.SubmitAddress)
	router.Get("/address", handler.GetAddresses)
	router.Post("/address/name", handler.NameAddress)
	router.Delete("/address", handler.DeleteAddress)

	router.Post("/clear", handler.ClearAddresses)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":  404,
			"error": "404: Not Found",
		})
	})

}
