package routes

import (
	"github.com/elanticrypt0/denario_go/pkg/features"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func recordsRoutes(appcore *webcore.AppCore, api fiber.Router) {
	records := api.Group("/records")
	records.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllRecords(c, appcore)
	})

	records.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneRecord(c, appcore)
	})

	records.Post("/", func(c *fiber.Ctx) error {
		return features.CreateRecord(c, appcore)
	})

	records.Post("/accountsmove/:year/:month", func(c *fiber.Ctx) error {
		return features.CreateAccountsMove(c, appcore)
	})

	records.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateRecord(c, appcore)
	})

	records.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteRecord(c, appcore)
	})

	records.Get("/:year/:month", func(c *fiber.Ctx) error {
		return features.FindRecordsByYearAndMonth(c, appcore)
	})
}
