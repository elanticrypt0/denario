package webcore_features

import (
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(appcore *webcore.AppCore) {

	// setup
	setup := appcore.Fiber.Group("/api/setup")
	if appcore.Config.App_setup_enabled {
		setup.Get("/", func(c *fiber.Ctx) error {
			// return webcore_features.Setup(c)
			return Setup(c, appcore)
		})
	}

	//status
	status := appcore.Fiber.Group("/api/status")
	status.Get("/", func(c *fiber.Ctx) error {
		return Status(c)
	})

	status.Get("/getenv", func(c *fiber.Ctx) error {
		return ReadEnv(c)
	})

	seeder := appcore.Fiber.Group("/api/seeder")
	if appcore.Config.App_setup_enabled {
		seeder.Get("/api/seed/", func(c *fiber.Ctx) error {
			return Seed(c)
		})
		seeder.Get("/api/seed/:table_name", func(c *fiber.Ctx) error {
			return Seed(c)
		})
	}
}
