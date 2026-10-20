package router

import (
	"os"
	"strings"

	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/handler"
	"Frank2006x/Vmate/internals/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func New(readQueries, writeQueries *db.Queries) *fiber.App {
	app := fiber.New()
	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(allowedOrigins, ","),
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	}))
	app.Use(logger.New())
	app.Get("/health", func(c fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) })
	api := app.Group("/api/v1")
	RegisterDepartmentRoutes(api, handler.NewDepartmentHandler(service.NewDepartmentService(readQueries, writeQueries)))
	RegisterStudentRoutes(api, handler.NewStudentHandler(service.NewStudentService(readQueries, writeQueries)))
	RegisterFacultyRoutes(api, handler.NewFacultyHandler(service.NewFacultyService(readQueries, writeQueries)))
	RegisterCourseRoutes(api, handler.NewCourseHandler(service.NewCourseService(readQueries, writeQueries)))
	return app
}
