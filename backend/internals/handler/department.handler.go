package handler

import (
	"Frank2006x/Vmate/internals/model"
	"Frank2006x/Vmate/internals/service"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

type DepartmentHandler struct{ service *service.DepartmentService }

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: s}
}
func (h *DepartmentHandler) List(c fiber.Ctx) error {
	items, err := h.service.List(c.Context())
	if err != nil {
		return respondError(c, 500, err)
	}
	return c.JSON(items)
}
func (h *DepartmentHandler) Get(c fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return respondError(c, 400, err)
	}
	item, err := h.service.Get(c.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		return respondError(c, 404, fmt.Errorf("department not found"))
	}
	if err != nil {
		return respondError(c, 500, err)
	}
	return c.JSON(item)
}
func (h *DepartmentHandler) Create(c fiber.Ctx) error {
	var input model.DepartmentRequest
	if err := c.Bind().JSON(&input); err != nil {
		return respondError(c, 400, err)
	}
	item, err := h.service.Create(c.Context(), input)
	if err != nil {
		return respondError(c, 400, err)
	}
	return c.Status(201).JSON(item)
}
func (h *DepartmentHandler) Update(c fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return respondError(c, 400, err)
	}
	var input model.DepartmentRequest
	if err = c.Bind().JSON(&input); err != nil {
		return respondError(c, 400, err)
	}
	item, err := h.service.Update(c.Context(), id, input)
	if err != nil {
		return respondError(c, 400, err)
	}
	return c.JSON(item)
}
func (h *DepartmentHandler) Delete(c fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return respondError(c, 400, err)
	}
	if err = h.service.Delete(c.Context(), id); err != nil {
		return respondError(c, 500, err)
	}
	return c.SendStatus(204)
}
