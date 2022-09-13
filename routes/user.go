package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/res0lution/digital-house/handlers"
)

func SetupUserRoutes(grp fiber.Router, handlers *handlers.Handler) {
	useRoute := grp.Group("/user")
	useRoute.Post("/register", handlers.UserRegister)
}