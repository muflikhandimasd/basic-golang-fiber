package controllers

import (
	"basic-golang-fiber/database"
	"basic-golang-fiber/models"
	"basic-golang-fiber/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Register creates a new user account
func Register(c *fiber.Ctx) error {
	var payload models.CreateUserSchema

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

	// Check if user already exists
	existingUser, _ := database.GetUserByEmail(payload.Email)
	if existingUser != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  false,
			"message": "User with this email already exists",
		})
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to hash password",
		})
	}

	// Set default role if not provided
	role := payload.Role
	if role == "" {
		role = "student"
	}

	// Create user
	user := &models.User{
		Email:    payload.Email,
		Password: hashedPassword,
		FullName: payload.FullName,
		Role:     role,
	}

	err = database.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  false,
				"message": "User with this email already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": true,
		"data": fiber.Map{
			"user": models.UserResponse{
				ID:        user.ID,
				Email:     user.Email,
				FullName:  user.FullName,
				Role:      user.Role,
				CreatedAt: user.CreatedAt,
			},
			"token": token,
		},
	})
}

// Login authenticates a user
func Login(c *fiber.Ctx) error {
	var payload models.LoginSchema

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

	// Get user by email
	user, err := database.GetUserByEmail(payload.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid email or password",
		})
	}

	// Check password
	if !utils.CheckPassword(payload.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid email or password",
		})
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data": fiber.Map{
			"user": models.UserResponse{
				ID:        user.ID,
				Email:     user.Email,
				FullName:  user.FullName,
				Role:      user.Role,
				CreatedAt: user.CreatedAt,
			},
			"token": token,
		},
	})
}

// GetProfile returns the current user's profile
func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int64)

	user, err := database.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
		"data": models.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FullName:  user.FullName,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
	})
}
