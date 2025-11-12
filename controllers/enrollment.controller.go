package controllers

import (
	"basic-golang-fiber/database"
	"basic-golang-fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// EnrollCourse enrolls a user in a course
func EnrollCourse(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int64)

	var payload models.CreateEnrollmentSchema
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
	_, err := database.GetCourseByID(payload.CourseID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Course not found",
		})
	}

	// Check if already enrolled
	existingEnrollment, _ := database.GetEnrollmentByUserAndCourse(userID, payload.CourseID)
	if existingEnrollment != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  false,
			"message": "You are already enrolled in this course",
		})
	}

	enrollment := &models.Enrollment{
		UserID:   userID,
		CourseID: payload.CourseID,
		Progress: 0,
	}

	err = database.CreateEnrollment(enrollment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Successfully enrolled in course",
		"data":    enrollment,
	})
}

// GetMyEnrollments retrieves all enrollments for logged-in user
func GetMyEnrollments(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int64)

	enrollments, err := database.GetEnrollmentsByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   enrollments,
	})
}

// GetEnrollmentsByCourse retrieves all enrollments for a course (instructor/admin only)
func GetEnrollmentsByCourse(c *fiber.Ctx) error {
	courseID, err := strconv.ParseInt(c.Params("courseId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid course ID",
		})
	}

	enrollments, err := database.GetEnrollmentsByCourse(courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   enrollments,
	})
}

// UpdateEnrollmentProgress updates progress for an enrollment
func UpdateEnrollmentProgress(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int64)
	enrollmentID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid enrollment ID",
		})
	}

	// Check if enrollment exists and belongs to user
	enrollment, err := database.GetEnrollmentByID(enrollmentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Enrollment not found",
		})
	}

	if enrollment.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "You are not authorized to update this enrollment",
		})
	}

	var payload struct {
		Progress int `json:"progress" validate:"required,min=0,max=100"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	err = database.UpdateEnrollmentProgress(enrollmentID, payload.Progress)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// Get updated enrollment
	updatedEnrollment, _ := database.GetEnrollmentByID(enrollmentID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   updatedEnrollment,
	})
}
