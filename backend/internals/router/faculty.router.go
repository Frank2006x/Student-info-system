package router

import (
	"Frank2006x/Vmate/internals/handler"
	"github.com/gofiber/fiber/v3"
)

func RegisterFacultyRoutes(api fiber.Router, h *handler.FacultyHandler) {
	route := api.Group("/faculty")
	route.Get("/", h.List)
	route.Post("/", h.Create)
	route.Get("/:id", h.Get)
	route.Patch("/:id", h.Update)
	route.Delete("/:id", h.Delete)
}
