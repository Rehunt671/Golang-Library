package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingsglaive/app/controllers"
	"github.com/kingsglaive/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	routes := a.Group("/api/v1")

	// Routes for POST method:
	routes.Post("/book", middleware.JWTProtected(), controllers.CreateBook)

	// Routes for PUT method:
	routes.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)

	// Routes for DELETE method with book ID as a parameter:
	routes.Delete("/book/:id", middleware.JWTProtected(), controllers.DeleteBook)
}
