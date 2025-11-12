package controllers

import (
	"basic-golang-fiber/database"
	"basic-golang-fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CreateCourse creates a new course
func CreateCourse(c *fiber.Ctx) error {
	var payload models.CreateCourseSchema

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

	// Verify instructor exists
	_, err := database.GetUserByID(payload.InstructorID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Instructor not found",
		})
	}

	// Set defaults
	level := payload.Level
	if level == "" {
		level = "beginner"
	}
	status := payload.Status
	if status == "" {
		status = "draft"
	}

	course := &models.Course{
		Title:        payload.Title,
		Description:  payload.Description,
		InstructorID: payload.InstructorID,
		CategoryID:   payload.CategoryID,
		Price:        payload.Price,
		Thumbnail:    payload.Thumbnail,
		Level:        level,
		Status:       status,
	}

	err = database.CreateCourse(course)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": true,
		"data":   course,
	})
}

// GetCourses retrieves all courses with pagination
func GetCourses(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)
	status := c.Query("status", "")

	courses, err := database.GetAllCourses(limit, offset, status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   courses,
	})
}

// GetCourseByID retrieves a specific course
func GetCourseByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid course ID",
		})
	}

	course, err := database.GetCourseWithDetails(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   course,
	})
}

// GetMyCourses retrieves courses by logged-in instructor
func GetMyCourses(c *fiber.Ctx) error {
	instructorID := c.Locals("userID").(int64)
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	courses, err := database.GetCoursesByInstructor(instructorID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   courses,
	})
}

// UpdateCourse updates a course
func UpdateCourse(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid course ID",
		})
	}

	var payload models.UpdateCourseSchema
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// Check if course exists
	course, err := database.GetCourseByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Course not found",
		})
	}

	// Check authorization (only course instructor or admin can update)
	userRole := c.Locals("userRole").(string)
	userID := c.Locals("userID").(int64)
	if userRole != "admin" && course.InstructorID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You are not authorized to update this course",
		})
	}

	// Build updates map
	updates := make(map[string]interface{})
	if payload.Title != "" {
		updates["title"] = payload.Title
	}
	if payload.Description != "" {
		updates["description"] = payload.Description
	}
	if payload.CategoryID != nil {
		updates["category_id"] = payload.CategoryID
	}
	if payload.Price > 0 {
		updates["price"] = payload.Price
	}
	if payload.Thumbnail != "" {
		updates["thumbnail"] = payload.Thumbnail
	}
	if payload.Level != "" {
		updates["level"] = payload.Level
	}
	if payload.Status != "" {
		updates["status"] = payload.Status
	}

	err = database.UpdateCourse(id, updates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// Get updated course
	updatedCourse, _ := database.GetCourseWithDetails(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   updatedCourse,
	})
}

// DeleteCourse deletes a course
func DeleteCourse(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid course ID",
		})
	}

	// Check if course exists
	course, err := database.GetCourseByID(id)
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
			"message": "You are not authorized to delete this course",
		})
	}

	err = database.DeleteCourse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Course deleted successfully",
	})
}
