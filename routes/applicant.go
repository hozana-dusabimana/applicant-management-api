package routes

import (
	"job-tracker/controllers"
	"job-tracker/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Initialize Redis connection
	controllers.InitRedis()
	
	// Setup applicant routes with middleware
	api := app.Group("/applicants")
	
	// Add request logging middleware
	api.Use(middleware.RequestLogger())
	
	// CRUD operations for applicants
	api.Post("/", controllers.CreateApplicant)
	api.Get("/", controllers.GetApplicants)
	api.Get("/:id", controllers.GetApplicant)
	api.Put("/:id", controllers.UpdateApplicant)
	api.Delete("/:id", controllers.DeleteApplicant)
}
