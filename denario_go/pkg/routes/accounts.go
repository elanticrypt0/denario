package routes

import (
	"github.com/elanticrypt0/denario_go/pkg/features"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func accountsRoutes(appcore *webcore.AppCore, api fiber.Router) {
	accounts := api.Group("/accounts")

	accounts.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllAccounts(c, appcore)
	})

	accounts.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneAccount(c, appcore)
	})

	accounts.Post("/", func(c *fiber.Ctx) error {
		return features.CreateAccount(c, appcore)
	})

	accounts.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateAccount(c, appcore)
	})

	accounts.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteAccount(c, appcore)
	})
}
