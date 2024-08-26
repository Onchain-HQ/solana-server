package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/Onchain-HQ/solana-server/pkg/api"
	handler "github.com/Onchain-HQ/solana-server/pkg/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	godotenv.Load()

	app := fiber.New(fiber.Config{
		// 1.5 GB
		BodyLimit: 1610612736,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if strings.Contains(err.Error(), "panic") {
				return c.Status(code).SendString("Internal Server Error")
			}

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			return c.Status(code).SendString(err.Error())
		},
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	/**
	db, err := database.New()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	*/

	handler := handler.New()
	var router fiber.Router = app
	api.New(router, handler)

	if err := app.Listen(":" + getenv("PORT", "4000")); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
