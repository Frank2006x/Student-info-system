package router

import (
	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/handler"
	"Frank2006x/Vmate/internals/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func New(queries *db.Queries) *fiber.App {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:        "[${time}] ${status} - ${latency} ${method} ${path} ${error}\n",
		TimeZone:      "UTC",
		DisableColors: true,
	}))
	app.Get("/health", func(c fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) })
	api := app.Group("/api/v1")
	RegisterDepartmentRoutes(api, handler.NewDepartmentHandler(service.NewDepartmentService(queries)))
	RegisterStudentRoutes(api, handler.NewStudentHandler(service.NewStudentService(queries)))
	RegisterFacultyRoutes(api, handler.NewFacultyHandler(service.NewFacultyService(queries)))
	RegisterCourseRoutes(api, handler.NewCourseHandler(service.NewCourseService(queries)))
	return app
}
