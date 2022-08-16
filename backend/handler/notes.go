package handler

import (
	"go-api/models"
	"go-api/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetNotes(c *fiber.Ctx) error {
	notes := &models.Notes{}

	log.Println(notes.FindNotes(c.Context()))

	return c.JSON(map[string]string{
		"notes": "Hello",
	})
}

func (h *Handler) CreateNotes(c *fiber.Ctx) error {
	notes := &models.Notes{}

	if err := c.BodyParser(notes); err != nil {
		return c.JSON(utils.ErrorResponse(err, "JSON payload invalid"))
	}

	updated, err := notes.SaveNotes(c.Context())
	if err != nil {
		return c.JSON(utils.ErrorResponse(err, "Unable to create a notes"))
	}

	return c.JSON(map[string]int64{
		"rows_affected": updated,
	})
}
