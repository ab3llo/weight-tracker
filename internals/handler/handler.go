package handler

import (
	"github.com/ab3llo/weight-tracker/database"
	"github.com/ab3llo/weight-tracker/internals/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateWeight(c *fiber.Ctx) error {
	db := database.DB
	weight := new(model.Weight)

	err := c.BodyParser(weight)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	weight.ID = uuid.New()
	err = db.Create(&weight).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create weight", "data": err})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created Weight", "data": weight})
}

func GetWeights(c *fiber.Ctx) error {
	db := database.DB
	var weights []model.Weight

	db.Find(&weights)

	if len(weights) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No weights present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Weights Found", "data": weights})
}

func GetWeight(c *fiber.Ctx) error {
	db := database.DB
	var weight model.Weight

	id := c.Params("weightId")

	db.Find(&weight, "id = ?", id)

	if weight.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No weight present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Weight Found", "data": weight})
}

func UpdateWeight(c *fiber.Ctx) error {
	db := database.DB
	var weight model.Weight

	id := c.Params("weightId")

	db.Find(&weight, "id = ?", id)

	if weight.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No weight present", "data": nil})
	}

	var updateWeight model.Weight

	err := c.BodyParser(&updateWeight)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	//edit weight
	weight.Weight = updateWeight.Weight
	db.Save(&weight)

	return c.JSON(fiber.Map{"status": "success", "message": "Weight found", "data": weight})
}

func DeleteWeight(c *fiber.Ctx) error {
	db := database.DB
	var weight model.Weight

	id := c.Params("weightId")

	db.Find(&weight, "id = ?", id)

	if weight.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Weight present"})
	}

	err := db.Delete(&weight, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}
