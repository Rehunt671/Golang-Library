package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	//Create routes group.
	route := a.Group("/swagger")

	// Routes for GET method:
	route.Get("*", swagger.HandlerDefault) // get one user by id
}
