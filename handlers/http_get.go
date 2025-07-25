package handlers

import "github.com/gofiber/fiber/v2"

func (s *HTTPServer) GetPotions(c *fiber.Ctx) error {
	return c.JSON(s.app.GetPotions())
}

func (s *HTTPServer) GetCauldrons(c *fiber.Ctx) error {
	return c.JSON(s.app.GetCauldrons())
}
