package handler

import (
	"github.com/Ankit-2205/put-call-ratio/pkg/optionchain"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service optionchain.Service
}

func (h *handler) GetOptionChain(c *fiber.Ctx) error {

	optionchain, err := h.service.GetOptionChain()
	if err != nil {
		return err
	}

	return c.JSON(optionchain)
}

func NewHandler(service optionchain.Service) Handler {
	return &handler{
		service: service,
	}
}
