package routes

import (
	"github.com/elanticrypt0/denario_go/pkg/features"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func categoriesRoutes(appcore *webcore.AppCore, api fiber.Router) {
	categories := api.Group("/categories")
	categories.Get("/", func(c *fiber.Ctx) error {
		return features.FindAllCategories(c, appcore)
	})
	categories.Get("/:id", func(c *fiber.Ctx) error {
		return features.FindOneCategory(c, appcore)
	})

	categories.Post("/", func(c *fiber.Ctx) error {
		return features.CreateCategory(c, appcore)
	})

	categories.Put("/:id", func(c *fiber.Ctx) error {
		return features.UpdateCategory(c, appcore)
	})

	categories.Delete("/:id", func(c *fiber.Ctx) error {
		return features.DeleteCategory(c, appcore)
	})
}
