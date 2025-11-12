package main

import (
	"log"

	"basic-golang-fiber/controllers"
	"basic-golang-fiber/database"
	"basic-golang-fiber/initializers"
	"basic-golang-fiber/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func init() {
	// Load configuration
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load environment variables:", err)
	}

	// Initialize PostgreSQL database
	dbConfig := database.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUserName,
		Password: config.DBUserPassword,
		DBName:   config.DBName,
		SSLMode:  config.DBSSLMode,
	}

	if err := database.InitDB(dbConfig); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run database migrations
	if err := database.RunMigrations(dbConfig); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
}

func main() {
	// Initialize template engine
	engine := html.New("./views", ".html")

	// Create Fiber app with template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,http://localhost:9000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PATCH, PUT, DELETE",
		AllowCredentials: true,
	}))

	// API Routes
	api := app.Group("/api")

	// Welcome route
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to eCourses API - Golang with PostgreSQL",
		})
	})

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Get("/profile", middleware.AuthMiddleware, controllers.GetProfile)

	// Category routes (public read, admin write)
	categories := api.Group("/categories")
	categories.Get("/", controllers.GetAllCategories)
	categories.Get("/:id", controllers.GetCategoryByID)
	categories.Post("/", middleware.AuthMiddleware, middleware.AdminMiddleware, controllers.CreateCategory)
	categories.Put("/:id", middleware.AuthMiddleware, middleware.AdminMiddleware, controllers.UpdateCategory)
	categories.Delete("/:id", middleware.AuthMiddleware, middleware.AdminMiddleware, controllers.DeleteCategory)

	// Course routes
	courses := api.Group("/courses")
	courses.Get("/", controllers.GetCourses)
	courses.Get("/:id", controllers.GetCourseByID)
	courses.Get("/:id/lessons", controllers.GetLessonsByCourse)
	courses.Get("/:courseId/enrollments", middleware.AuthMiddleware, middleware.InstructorOrAdminMiddleware, controllers.GetEnrollmentsByCourse)

	// Protected course routes (instructor/admin)
	courses.Post("/", middleware.AuthMiddleware, middleware.InstructorOrAdminMiddleware, controllers.CreateCourse)
	courses.Get("/my/courses", middleware.AuthMiddleware, middleware.InstructorOrAdminMiddleware, controllers.GetMyCourses)
	courses.Put("/:id", middleware.AuthMiddleware, middleware.InstructorOrAdminMiddleware, controllers.UpdateCourse)
	courses.Delete("/:id", middleware.AuthMiddleware, middleware.InstructorOrAdminMiddleware, controllers.DeleteCourse)

	// Lesson routes
	lessons := api.Group("/lessons")
	lessons.Get("/:id", controllers.GetLessonByID)
	lessons.Post("/", middleware.AuthMiddleware, middleware.InstructorOrAdminMiddleware, controllers.CreateLesson)
	lessons.Put("/:id", middleware.AuthMiddleware, middleware.InstructorOrAdminMiddleware, controllers.UpdateLesson)
	lessons.Delete("/:id", middleware.AuthMiddleware, middleware.InstructorOrAdminMiddleware, controllers.DeleteLesson)

	// Enrollment routes
	enrollments := api.Group("/enrollments")
	enrollments.Post("/", middleware.AuthMiddleware, controllers.EnrollCourse)
	enrollments.Get("/my", middleware.AuthMiddleware, controllers.GetMyEnrollments)
	enrollments.Put("/:id/progress", middleware.AuthMiddleware, controllers.UpdateEnrollmentProgress)

	// Admin Panel Routes
	admin := app.Group("/admin")

	// For now, admin panel is accessible without authentication for easy testing
	// In production, you should add admin authentication middleware
	admin.Get("/", controllers.AdminDashboard)
	admin.Get("/courses", controllers.AdminCourses)
	admin.Get("/users", controllers.AdminUsers)

	// Home route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/admin")
	})

	log.Println("🚀 Server starting on http://localhost:9000")
	log.Println("📊 Admin Panel: http://localhost:9000/admin")
	log.Println("🔌 API Endpoint: http://localhost:9000/api")
	log.Fatal(app.Listen(":9000"))
}
