package middleware

import (
	"basic-golang-fiber/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware validates JWT token
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Missing authorization header",
		})
	}

	// Extract token from "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid authorization header format",
		})
	}

	token := parts[1]
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid or expired token",
		})
	}

	// Store user info in context
	c.Locals("userID", claims.UserID)
	c.Locals("userEmail", claims.Email)
	c.Locals("userRole", claims.Role)

	return c.Next()
}

// AdminMiddleware checks if user is admin
func AdminMiddleware(c *fiber.Ctx) error {
	role := c.Locals("userRole")
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "Admin access required",
		})
	}
	return c.Next()
}

// InstructorOrAdminMiddleware checks if user is instructor or admin
func InstructorOrAdminMiddleware(c *fiber.Ctx) error {
	role := c.Locals("userRole")
	if role != "admin" && role != "instructor" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  false,
			"message": "Instructor or admin access required",
		})
	}
	return c.Next()
}
