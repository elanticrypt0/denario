package features

import (
	"strconv"

	"github.com/elanticrypt0/denario_go/pkg/models"
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

//! UPDATE!
// Ahora credits se sigue llamando así, pero tiene una función diferente: Ahora es para gastos "fijos" o constantes, como lo son los servicios, alquileres, etc.
// Ahora la tabla de credits tiene una columna "is_fixed" que si es true significa que es un gasto fijo, y si es false significa que es un gasto variable.
// Otro cambio importante es que si un gasto es "Fijo" al momento de crear los records relacionados esos records pueden ser mutables, para adaptar el gasto a la realidad.

func FindOneCredit(c *fiber.Ctx, appcore *webcore.AppCore) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.FindOneCredit(appcore, id))
}

func FindAllCredits(c *fiber.Ctx, appcore *webcore.AppCore) error {
	return c.JSON(models.FindAllCredits(appcore))
}

func CreateCredit(c *fiber.Ctx, appcore *webcore.AppCore) error {

	new_c := new(models.Credit)
	c.BodyParser(&new_c)
	credit := models.CreateCredit(appcore, new_c.Name, new_c.Comment, new_c.Amount, new_c.Payments, new_c.StartedAt, new_c.CategoryID, new_c.IsFixed)

	return c.JSON(credit)
}

func UpdateCredit(c *fiber.Ctx, appcore *webcore.AppCore) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	credit := models.FindOneCredit(appcore, id)
	c.BodyParser(&credit)
	credit = models.UpdateCredit(appcore, credit)

	return c.JSON(credit)
}

func DeleteCredit(c *fiber.Ctx, appcore *webcore.AppCore) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	return c.JSON(models.DeleteCredit(appcore, id))
}
