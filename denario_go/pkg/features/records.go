package features

import (
	"fmt"
	"strconv"

	"github.com/elanticrypt0/denario_go/pkg/models"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/elanticrypt0/denario_go/pkg/webcore/helpers"
	"github.com/gofiber/fiber/v2"
)

// Find a record by ID
// Return the record
// Return an error if there was a problem
func FindOneRecord(c *fiber.Ctx, appcore *webcore.AppCore) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.FindOneRecord(appcore, id))
}

// Find all records
// Return the records
// Return an error if there was a problem
func FindAllRecords(c *fiber.Ctx, appcore *webcore.AppCore) error {
	return c.JSON(models.FindAllRecords(appcore))
}

// Create a new record
// Return the new record
// Return an error if there was a problem
func CreateRecord(c *fiber.Ctx, appcore *webcore.AppCore) error {
	r := new(models.Record)
	c.BodyParser(&r)
	record := models.CreateRecord(appcore, r.Name, r.Amount, r.AmountIo, r.Comment, r.RecordDate, r.CategoryID, r.IsMutable, r.AccountID)
	return c.JSON(record)
}

// Crea una transferencia entre cuentas
// Return the new record
// Return an error if there was a problem
func CreateAccountsMove(c *fiber.Ctx, appcore *webcore.AppCore) error {

	month := c.Params("month", "0")
	year := c.Params("year", "0")
	month_str := helpers.AddZeroBeforeNumberFromStr(month)

	// ID de la categoría de transferencia
	transference_cat_id := 28

	r := new(models.Transference)
	c.BodyParser(&r)

	account_sender := models.FindOneAccount(appcore, r.AccountIDSender)
	account_reciber := models.FindOneAccount(appcore, r.AccountIDReciber)

	// datos del nuevo record (Previo a la transfrencia)
	// obtiene el total de una cuenta en previo a realizar la transfrencia
	account_sender_balance := models.GetTotalAmountOfAccount(appcore, year, month_str, int(account_sender.ID))
	account_reciber_balance := models.GetTotalAmountOfAccount(appcore, year, month_str, int(account_reciber.ID))

	// Calcula los valores previos y posteriores del sender
	account_sender_prev_amount := account_sender_balance.Total
	account_sender_post_amount := account_sender_balance.Total - r.Amount

	// Calcula los valores previos y posteriores del reciber
	account_reciber_prev_amount := account_reciber_balance.Total
	account_reciber_post_amount := account_reciber_balance.Total + r.Amount

	record_name := fmt.Sprintf("Desde %s a %s", account_sender.Short, account_reciber.Short)
	comment_prev := fmt.Sprintf("Prev.: %s: $%f | %s: $%f", account_sender.Short, account_sender_prev_amount, account_reciber.Short, account_reciber_prev_amount)

	models.CreateRecord(appcore, record_name, r.Amount, "out", comment_prev, r.RecordDate, transference_cat_id, false, r.AccountIDSender)

	// datos del nuevo record (Actual)

	comment_current := fmt.Sprintf("Actual: %s: $%f | %s: $%f", account_sender.Short, account_sender_post_amount, account_reciber.Short, account_reciber_post_amount)

	record := models.CreateRecord(appcore, record_name, r.Amount, "in", comment_current, r.RecordDate, transference_cat_id, false, r.AccountIDReciber)

	return c.JSON(record)
}

// Update a record
// Return the updated record
// Return an error if there was a problem
func UpdateRecord(c *fiber.Ctx, appcore *webcore.AppCore) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	record := models.FindOneRecord(appcore, id)
	c.BodyParser(&record)
	models.UpdateRecord(appcore, record)
	return c.JSON(record)
}

// Delete a record
// Return the deleted record
// Return an error if there was a problem
func DeleteRecord(c *fiber.Ctx, appcore *webcore.AppCore) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.DeleteRecord(appcore, id))
}

// find by Month and Year

func FindRecordsByYearAndMonth(c *fiber.Ctx, appcore *webcore.AppCore) error {
	month := c.Params("month", "0")
	year := c.Params("year", "0")
	month_str := helpers.AddZeroBeforeNumberFromStr(month)

	// records := models.FindRecordsByYearAndMonth(year, month_str, "")
	records := models.FindRecordsByYearAndMonth(appcore, year, month_str, "")

	return c.JSON(records)
}
