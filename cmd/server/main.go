package main

import (
	"github.com/ab3llo/weight-tracker/config"
	"github.com/ab3llo/weight-tracker/database"
	"github.com/ab3llo/weight-tracker/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()
	cfg := config.New()

	// connect to Database
	database.ConnectDB(cfg)

	// Set up routes
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":" + cfg.Port)
}
