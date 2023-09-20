package routes

import (
	"github.com/elanticrypt0/denario_go/pkg/features"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func creditsRoutes(appcore *webcore.AppCore, api fiber.Router) {
	credits := api.Group("/credits")
	credits.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllCredits(c, appcore)
	})

	credits.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneCredit(c, appcore)
	})

	credits.Post("/:category_id", func(c *fiber.Ctx) error {
		return features.CreateCredit(c, appcore)
	})

	credits.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateCredit(c, appcore)
	})

	credits.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteCredit(c, appcore)
	})
}
