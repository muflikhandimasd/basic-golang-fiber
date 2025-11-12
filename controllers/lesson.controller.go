package controllers

import (
	"basic-golang-fiber/database"
	"basic-golang-fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreateLesson creates a new lesson
func CreateLesson(c *fiber.Ctx) error {
	var payload models.CreateLessonSchema

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

	// Verify course exists
	course, err := database.GetCourseByID(payload.CourseID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Course not found",
		})
	}

	// Check authorization
	userRole := c.Locals("userRole").(string)
	userID := c.Locals("userID").(int64)
	if userRole != "admin" && course.InstructorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You are not authorized to add lessons to this course",
		})
	}

	lesson := &models.Lesson{
		CourseID:   payload.CourseID,
		Title:      payload.Title,
		Content:    payload.Content,
		VideoURL:   payload.VideoURL,
		Duration:   payload.Duration,
		OrderIndex: payload.OrderIndex,
		IsFree:     payload.IsFree,
	}

	err = database.CreateLesson(lesson)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": true,
		"data":   lesson,
	})
}

// GetLessonsByCourse retrieves all lessons for a course
func GetLessonsByCourse(c *fiber.Ctx) error {
	courseID, err := strconv.ParseInt(c.Params("courseId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid course ID",
		})
	}

	lessons, err := database.GetLessonsByCourse(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   lessons,
	})
}

// GetLessonByID retrieves a specific lesson
func GetLessonByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid lesson ID",
		})
	}

	lesson, err := database.GetLessonByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   lesson,
	})
}

// UpdateLesson updates a lesson
func UpdateLesson(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid lesson ID",
		})
	}

	var payload models.UpdateLessonSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// Check if lesson exists
	lesson, err := database.GetLessonByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Lesson not found",
		})
	}

	// Get course to check authorization
	course, err := database.GetCourseByID(lesson.CourseID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Course not found",
		})
	}

	// Check authorization
	userRole := c.Locals("userRole").(string)
	userID := c.Locals("userID").(int64)
	if userRole != "admin" && course.InstructorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You are not authorized to update this lesson",
		})
	}

	// Build updates map
	updates := make(map[string]interface{})
	if payload.Title != "" {
		updates["title"] = payload.Title
	}
	if payload.Content != "" {
		updates["content"] = payload.Content
	}
	if payload.VideoURL != "" {
		updates["video_url"] = payload.VideoURL
	}
	if payload.Duration > 0 {
		updates["duration"] = payload.Duration
	}
	if payload.OrderIndex > 0 {
		updates["order_index"] = payload.OrderIndex
	}
	if payload.IsFree != nil {
		updates["is_free"] = *payload.IsFree
	}

	err = database.UpdateLesson(id, updates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// Get updated lesson
	updatedLesson, _ := database.GetLessonByID(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   updatedLesson,
	})
}

// DeleteLesson deletes a lesson
func DeleteLesson(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid lesson ID",
		})
	}

	// Check if lesson exists
	lesson, err := database.GetLessonByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Lesson not found",
		})
	}

	// Get course to check authorization
	course, err := database.GetCourseByID(lesson.CourseID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Course not found",
		})
	}

	// Check authorization
	userRole := c.Locals("userRole").(string)
	userID := c.Locals("userID").(int64)
	if userRole != "admin" && course.InstructorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You are not authorized to delete this lesson",
		})
	}

	err = database.DeleteLesson(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Lesson deleted successfully",
	})
}
