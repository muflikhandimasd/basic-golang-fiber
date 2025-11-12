package controllers

import (
	"basic-golang-fiber/database"
	"basic-golang-fiber/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// CreateCategory creates a new category (admin only)
func CreateCategory(c *fiber.Ctx) error {
	var payload models.CreateCategorySchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	category := &models.Category{
		Name:        payload.Name,
		Description: payload.Description,
		Slug:        payload.Slug,
	}

	err := database.CreateCategory(category)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  false,
				"message": "Category with this name or slug already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": true,
		"data":   category,
	})
}

// GetAllCategories retrieves all categories
func GetAllCategories(c *fiber.Ctx) error {
	categories, err := database.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   categories,
	})
}

// GetCategoryByID retrieves a specific category
func GetCategoryByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid category ID",
		})
	}

	category, err := database.GetCategoryByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   category,
	})
}

// UpdateCategory updates a category (admin only)
func UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid category ID",
		})
	}

	var payload models.UpdateCategorySchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// Build updates map
	updates := make(map[string]interface{})
	if payload.Name != "" {
		updates["name"] = payload.Name
	}
	if payload.Description != "" {
		updates["description"] = payload.Description
	}
	if payload.Slug != "" {
		updates["slug"] = payload.Slug
	}

	err = database.UpdateCategory(id, updates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// Get updated category
	updatedCategory, _ := database.GetCategoryByID(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   updatedCategory,
	})
}

// DeleteCategory deletes a category (admin only)
func DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid category ID",
		})
	}

	err = database.DeleteCategory(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Category deleted successfully",
	})
}
