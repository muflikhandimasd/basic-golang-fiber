package controllers

import (
	"basic-golang-fiber/database"

	"github.com/gofiber/fiber/v2"
)

// AdminDashboard renders the admin dashboard
func AdminDashboard(c *fiber.Ctx) error {
	// Get statistics
	totalCourses, _ := database.CountCourses()
	totalUsers, _ := database.CountUsers()

	// Get published courses count
	publishedCourses, _ := database.GetAllCourses(1000, 0, "published")

	// Count enrollments
	var totalEnrollments int
	query := "SELECT COUNT(*) FROM enrollments"
	database.DB.QueryRow(query).Scan(&totalEnrollments)

	return c.Render("admin/dashboard", fiber.Map{
		"TotalCourses":     totalCourses,
		"TotalUsers":       totalUsers,
		"TotalEnrollments": totalEnrollments,
		"PublishedCourses": len(publishedCourses),
	})
}

// AdminCourses renders the courses management page
func AdminCourses(c *fiber.Ctx) error {
	courses, err := database.GetAllCourses(100, 0, "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading courses")
	}

	return c.Render("admin/courses", fiber.Map{
		"Courses": courses,
	})
}

// AdminUsers renders the users management page
func AdminUsers(c *fiber.Ctx) error {
	users, err := database.GetAllUsers(100, 0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading users")
	}

	return c.Render("admin/users", fiber.Map{
		"Users": users,
	})
}
