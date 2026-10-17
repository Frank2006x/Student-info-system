package router

import (
	"Frank2006x/Vmate/internals/handler"
	"github.com/gofiber/fiber/v3"
)

func RegisterCourseRoutes(api fiber.Router, h *handler.CourseHandler) {
	route := api.Group("/courses")
	route.Get("/", h.List)
	route.Post("/", h.Create)
	route.Get("/:id", h.Get)
	route.Patch("/:id", h.Update)
	route.Delete("/:id", h.Delete)
}
