package features

import (
	"strconv"

	"github.com/elanticrypt0/denario_go/pkg/models"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/elanticrypt0/denario_go/pkg/webcore/helpers"
	"github.com/gofiber/fiber/v2"
)

// Find outcomes of a date YEAR/MONTH
// Return the record
// Return an error if there was a problem
func FindBalancesOfDate(c *fiber.Ctx, appcore *webcore.AppCore) error {
	month := c.Params("month", "0")
	year := c.Params("year", "0")
	amount_io := c.Params("amount_io", "in")

	month_str := helpers.AddZeroBeforeNumberFromStr(month)

	balance := models.GetTotalAmountOfDateByIo(appcore, year, month_str, amount_io)

	return c.JSON(balance)
}

// Find outcomes of a date YEAR/MONTH
// Return the record
// Return an error if there was a problem
func FindBalancesFixedOfDate(c *fiber.Ctx, appcore *webcore.AppCore) error {
	month := c.Params("month", "0")
	year := c.Params("year", "0")
	amount_io := c.Params("amount_io", "in")

	month_str := helpers.AddZeroBeforeNumberFromStr(month)

	balance := models.GetTotalAmountOfDateByIoFixed(appcore, year, month_str, amount_io)

	return c.JSON(balance)
}

// Find outcomes of a date YEAR/MONTH
// Return the record
// Return an error if there was a problem
func FindBalancesCreditsOfDate(c *fiber.Ctx, appcore *webcore.AppCore) error {
	month := c.Params("month", "0")
	year := c.Params("year", "0")
	amount_io := c.Params("amount_io", "in")

	month_str := helpers.AddZeroBeforeNumberFromStr(month)

	balance := models.GetTotalAmountOfDateByIoCredits(appcore, year, month_str, amount_io)

	return c.JSON(balance)
}

// Find outcomes of a date YEAR/MONTH
// Return the record
// Return an error if there was a problem
func FindBalancesSavingsOfDate(c *fiber.Ctx, appcore *webcore.AppCore) error {
	month := c.Params("month", "0")
	year := c.Params("year", "0")

	month_str := helpers.AddZeroBeforeNumberFromStr(month)

	balance := models.GetTotalAmountOfDateByIoSavings(appcore, year, month_str)

	return c.JSON(balance)
}

// Find outcomes of a date YEAR/MONTH
// Return the record
// Return an error if there was a problem
func GetTotalAmountOfAccount(c *fiber.Ctx, appcore *webcore.AppCore) error {
	account_id, _ := strconv.Atoi(c.Params("account_id", "0"))
	month := c.Params("month", "0")
	year := c.Params("year", "0")

	month_str := helpers.AddZeroBeforeNumberFromStr(month)

	balance := models.GetTotalAmountOfAccount(appcore, year, month_str, account_id)

	return c.JSON(balance)
}
