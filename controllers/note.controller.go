package controllers

import (
	"fmt"
	"strings"
	"time"

	"basic-golang-fiber/initializers"
	"basic-golang-fiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateNoteHandler(c *fiber.Ctx) error {
	var payload *models.CreateNoteSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newNote := models.Note{
		Title:     payload.Title,
		Content:   payload.Content,
		Category:  payload.Category,
		Published: payload.Published,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := initializers.DB.Create(&newNote)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": false, "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": true, "data": fiber.Map{"note": newNote}})
}

func FindNotes(c *fiber.Ctx) error {

	var notes []models.Note
	results := initializers.DB.Order("id DESC").Find(&notes)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "message": "Success Get Notes", "data": notes})
}

func UpdateNote(c *fiber.Ctx) error {
	noteId := c.Params("noteId")

	var payload *models.UpdateNoteSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	var note models.Note
	result := initializers.DB.First(&note, "id = ?", noteId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Title != "" {
		updates["title"] = payload.Title
	}
	if payload.Category != "" {
		updates["category"] = payload.Category
	}
	if payload.Content != "" {
		updates["content"] = payload.Content
	}

	if payload.Published != nil {
		updates["published"] = payload.Published
	}

	updates["updated_at"] = time.Now()

	initializers.DB.Model(&note).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "data": note})
}

func FindNoteById(c *fiber.Ctx) error {
	noteId := c.Params("noteId")

	var note models.Note
	result := initializers.DB.First(&note, "id = ?", noteId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": false, "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	noteId := c.Params("noteId")

	result := initializers.DB.Delete(&models.Note{}, "id = ?", noteId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": false, "message": "No note with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Note Successfully Deleted",
	})
}

func FakerNotes(c *fiber.Ctx) error {
	total := c.QueryInt("total", 1)
	for i := 0; i < total; i++ {
		now := time.Now()
		newNote := models.Note{
			Title:     fmt.Sprintf("Title %d", i+15500),
			Content:   fmt.Sprintf("Content %d", i),
			Category:  fmt.Sprintf("Category %d", i),
			Published: true,
			CreatedAt: now,
			UpdatedAt: now,
		}

		initializers.DB.Create(&newNote)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "message": "Success Create Notes"})
}
