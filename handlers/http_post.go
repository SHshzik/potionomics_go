package handlers

import (
	"net/http"

	"github.com/SHshzik/potionomics_go/domain"
	"github.com/gofiber/fiber/v2"
)

type generateForm struct {
	PotionID   string `json:"potion_id" validate:"required"`
	CauldronID string `json:"cauldron_id" validate:"required"`
}

func (s *HTTPServer) Generate(c *fiber.Ctx) error {
	formGenerate := generateForm{}

	err := c.BodyParser(&formGenerate)
	if err != nil {
		s.l.Error(err, "http - generate - body parser")

		return errorResponse(c, http.StatusBadRequest, "Bad generate params")
	}

	err = s.v.Struct(formGenerate)
	if err != nil {
		s.l.Error(err, "http - generate - validate")

		return errorResponse(c, http.StatusUnprocessableEntity, err.Error())
	}

	potion, err := s.app.GetPotion(formGenerate.PotionID)
	if err != nil {
		s.l.Error(err, "http - generate - potion not found")
		return errorResponse(c, http.StatusNotFound, "Potion not found")
	}

	cauldron, err := s.app.GetCauldron(formGenerate.CauldronID)
	if err != nil {
		s.l.Error(err, "http - generate - cauldron not found")
		return errorResponse(c, http.StatusNotFound, "Cauldron not found")
	}

	return c.JSON(s.app.Generate(domain.GenerateRequest{
		Potion:   potion,
		Cauldron: cauldron,
	}))
}
