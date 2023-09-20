package routes

import (
	"github.com/elanticrypt0/denario_go/pkg/features"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func balancesRoutes(appcore *webcore.AppCore, api fiber.Router) {
	balances := api.Group("/balances")
	balances.Get("/:amount_io/:year/:month", func(c *fiber.Ctx) error {
		return features.FindBalancesOfDate(c, appcore)
	})
	balances.Get("/:amount_io/:year/:month/fixed", func(c *fiber.Ctx) error {
		return features.FindBalancesFixedOfDate(c, appcore)
	})
	balances.Get("/:amount_io/:year/:month/credits", func(c *fiber.Ctx) error {
		return features.FindBalancesCreditsOfDate(c, appcore)
	})
	balances.Get("/savings/:year/:month", func(c *fiber.Ctx) error {
		return features.FindBalancesSavingsOfDate(c, appcore)
	})
	balances.Get("/account/:account_id/:year/:month", func(c *fiber.Ctx) error {
		return features.GetTotalAmountOfAccount(c, appcore)
	})
}
