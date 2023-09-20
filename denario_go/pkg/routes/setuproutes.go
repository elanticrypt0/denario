package routes

import (
	"github.com/elanticrypt0/denario_go/pkg/webcore"
)

func SetupFeaturesRoutes(appcore *webcore.AppCore) {
	api := appcore.Fiber.Group("/api")
	// categories
	categoriesRoutes(appcore, api)
	// credits
	creditsRoutes(appcore, api)
	// accounts
	accountsRoutes(appcore, api)
	// records
	recordsRoutes(appcore, api)
	// balances
	balancesRoutes(appcore, api)
}
