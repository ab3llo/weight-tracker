package routes

import (
	"github.com/ab3llo/weight-tracker/internals/handler"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(router fiber.Router) {
	weight := router.Group("/weight")
	weight.Post("/", handler.CreateWeight)
	weight.Get("/", handler.GetWeights)
	weight.Get("/:weightId", handler.GetWeight)
	weight.Put("/:weightId", handler.UpdateWeight)
	weight.Delete("/:weightId", handler.DeleteWeight)
}
