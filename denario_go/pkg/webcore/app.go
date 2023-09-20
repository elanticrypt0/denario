package webcore

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppCore struct {
	Fiber  *fiber.App
	Db     *gorm.DB
	Config AppConfig
}
