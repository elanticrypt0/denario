package main

import (
	"log"

	"github.com/elanticrypt0/denario_go/pkg/routes"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/elanticrypt0/denario_go/pkg/webcore_features"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app_config := webcore.LoadConfig()

	appcore := webcore.AppCore{
		Fiber:  fiber.New(fiber.Config{}),
		Config: app_config,
		Db:     webcore.Connect2DB(&app_config),
	}

	// CORS
	// necesario para poder utilizar la WUI
	appcore.Fiber.Use(cors.New())
	appcore.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,http://localhost:65065",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	appcore.Fiber.Use(recover.New())

	// features routes
	routes.SetupFeaturesRoutes(&appcore)
	// webcore features
	webcore_features.SetupRoutes(&appcore)
	// static routes
	webcore.SetupStaticRoutes(appcore.Fiber)

	log.Fatal(appcore.Fiber.Listen(":"+appcore.Config.App_server_port), "Server is running on port "+appcore.Config.App_server_port)

}
