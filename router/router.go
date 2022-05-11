package router

import (
	"fmt"

	routes "github.com/ab3llo/weight-tracker/internals/routes/weight"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	fmt.Println("set up routes")
	api := app.Group("api", logger.New())
	routes.SetUpRoutes(api)
	fmt.Println("set up routes")
}
