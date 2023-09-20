package webcore_features

import (
	"github.com/elanticrypt0/denario_go/pkg/models"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func Setup(c *fiber.Ctx, appcore *webcore.AppCore) error {
	migrateModels(appcore)
	return c.SendString("Setup enabled. Models Migrated.")
}

func migrateModels(appcore *webcore.AppCore) {
	appcore.Db.AutoMigrate(&models.Category{}, &models.Credit{}, &models.Record{}, &models.Account{})
}
