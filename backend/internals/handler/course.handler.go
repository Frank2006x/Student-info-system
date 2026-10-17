package handler

import (
	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/model"
	"Frank2006x/Vmate/internals/service"

	"github.com/gofiber/fiber/v3"
)

type CourseHandler struct{ service *service.CourseService }

func NewCourseHandler(s *service.CourseService) *CourseHandler { return &CourseHandler{service: s} }
func (h *CourseHandler) List(c fiber.Ctx) error {
	items, err := h.service.List(c.Context(), db.ListCoursesParams{PageOffset: 0, PageSize: 100})
	if err != nil {
		return respondError(c, 500, err)
	}
	return c.JSON(items)
}
func (h *CourseHandler) Get(c fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return respondError(c, 400, err)
	}
	item, err := h.service.Get(c.Context(), id)
	if err != nil {
		return respondError(c, 404, err)
	}
	return c.JSON(item)
}
func (h *CourseHandler) Create(c fiber.Ctx) error {
	var input model.CourseRequest
	if err := c.Bind().JSON(&input); err != nil {
		return respondError(c, 400, err)
	}
	item, err := h.service.Create(c.Context(), input)
	if err != nil {
		return respondError(c, 400, err)
	}
	return c.Status(201).JSON(item)
}
func (h *CourseHandler) Update(c fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return respondError(c, 400, err)
	}
	var input model.CourseRequest
	if err = c.Bind().JSON(&input); err != nil {
		return respondError(c, 400, err)
	}
	item, err := h.service.Update(c.Context(), id, input)
	if err != nil {
		return respondError(c, 400, err)
	}
	return c.JSON(item)
}
func (h *CourseHandler) Delete(c fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return respondError(c, 400, err)
	}
	if err = h.service.Delete(c.Context(), id); err != nil {
		return respondError(c, 500, err)
	}
	return c.SendStatus(204)
}
