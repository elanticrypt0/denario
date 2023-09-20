package features

import (
	"strconv"

	"github.com/elanticrypt0/denario_go/pkg/models"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

// Find a Account by ID
// Return the Account
// Return an error if there was a problem
func FindOneAccount(c *fiber.Ctx, appcore *webcore.AppCore) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.FindOneAccount(appcore, id))
}

// Find all Accounts
// Return the Accounts
// Return an error if there was a problem
func FindAllAccounts(c *fiber.Ctx, appcore *webcore.AppCore) error {
	return c.JSON(models.FindAllAccounts(appcore))
}

// Create a new Account
// Return the new Account
// Return an error if there was a problem
func CreateAccount(c *fiber.Ctx, appcore *webcore.AppCore) error {
	r := new(models.Account)
	c.BodyParser(&r)
	Account := models.CreateAccount(appcore, r.Name, r.Short, r.Detail)
	return c.JSON(Account)
}

// Update a Account
// Return the updated Account
// Return an error if there was a problem
func UpdateAccount(c *fiber.Ctx, appcore *webcore.AppCore) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	Account := models.FindOneAccount(appcore, id)
	c.BodyParser(&Account)
	models.UpdateAccount(appcore, Account)
	return c.JSON(Account)
}

// Delete a Account
// Return the deleted Account
// Return an error if there was a problem
func DeleteAccount(c *fiber.Ctx, appcore *webcore.AppCore) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.DeleteAccount(appcore, id))
}
