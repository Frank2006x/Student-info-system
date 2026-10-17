package router

import (
	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/handler"
	"Frank2006x/Vmate/internals/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func New(queries *db.Queries) *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	}))
	app.Use(logger.New())
	app.Get("/health", func(c fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) })
	api := app.Group("/api/v1")
	RegisterDepartmentRoutes(api, handler.NewDepartmentHandler(service.NewDepartmentService(queries)))
	RegisterStudentRoutes(api, handler.NewStudentHandler(service.NewStudentService(queries)))
	RegisterFacultyRoutes(api, handler.NewFacultyHandler(service.NewFacultyService(queries)))
	RegisterCourseRoutes(api, handler.NewCourseHandler(service.NewCourseService(queries)))
	return app
}
