package handler

import (
	db "Frank2006x/Vmate/db/sqlc"
	"Frank2006x/Vmate/internals/model"
	"Frank2006x/Vmate/internals/service"

	"github.com/gofiber/fiber/v3"
)

type FacultyHandler struct{ service *service.FacultyService }

func NewFacultyHandler(s *service.FacultyService) *FacultyHandler { return &FacultyHandler{service: s} }
func (h *FacultyHandler) List(c fiber.Ctx) error {
	items, err := h.service.List(c.Context(), db.ListFacultyParams{PageOffset: 0, PageSize: 100})
	if err != nil {
		return respondError(c, 500, err)
	}
	return c.JSON(items)
}
func (h *FacultyHandler) Get(c fiber.Ctx) error {
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
func (h *FacultyHandler) Create(c fiber.Ctx) error {
	var input model.FacultyRequest
	if err := c.Bind().JSON(&input); err != nil {
		return respondError(c, 400, err)
	}
	item, err := h.service.Create(c.Context(), input)
	if err != nil {
		return respondError(c, 400, err)
	}
	return c.Status(201).JSON(item)
}
func (h *FacultyHandler) Update(c fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return respondError(c, 400, err)
	}
	var input model.FacultyRequest
	if err = c.Bind().JSON(&input); err != nil {
		return respondError(c, 400, err)
	}
	item, err := h.service.Update(c.Context(), id, input)
	if err != nil {
		return respondError(c, 400, err)
	}
	return c.JSON(item)
}
func (h *FacultyHandler) Delete(c fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return respondError(c, 400, err)
	}
	if err = h.service.Delete(c.Context(), id); err != nil {
		return respondError(c, 500, err)
	}
	return c.SendStatus(204)
}
