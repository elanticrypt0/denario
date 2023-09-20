package models

import (
	"strconv"

	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type Credit struct {
	gorm.Model
	Name       string   `json:"name"`
	Comment    string   `json:"comment"`
	Amount     float64  `json:"amount"`
	Payments   int      `json:"payments"`
	StartedAt  string   `json:"started_at"`
	FinishedAt string   `json:"finished_at"`
	CategoryID int      `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID"`
	IsFixed    bool     `json:"is_fixed"`
}

func FindOneCredit(appcore *webcore.AppCore, id int) Credit {
	var credit Credit
	appcore.Db.Preload("Category").First(&credit, id)
	return credit
}

func FindAllCredits(appcore *webcore.AppCore) []Credit {
	var credits []Credit
	appcore.Db.Preload("Category").Find(&credits)
	return credits
}

func CreateCredit(appcore *webcore.AppCore, name string, comment string, amount float64, payments int, started_at string, category_id int, isFixed bool) Credit {

	finished_at := getFinishAtDate(started_at, payments)

	credit := Credit{
		Name:       name,
		Comment:    comment,
		Amount:     amount,
		Payments:   payments,
		StartedAt:  started_at,
		FinishedAt: finished_at,
		CategoryID: category_id,
		IsFixed:    isFixed,
	}
	appcore.Db.Create(&credit)
	CreateRecordsForCredit(appcore, credit)
	return credit
}

func UpdateCredit(appcore *webcore.AppCore, credit Credit) Credit {

	credit.FinishedAt = getFinishAtDate(credit.StartedAt, credit.Payments)

	appcore.Db.Save(&credit)
	deleteAllRecordsOfThisCredit(appcore, int(credit.ID), credit.IsFixed)
	CreateRecordsForCredit(appcore, credit)
	return credit
}

func DeleteCredit(appcore *webcore.AppCore, id int) Credit {
	var credit Credit
	appcore.Db.First(&credit, id)
	// also remote all records associated with this credit
	deleteAllRecordsOfThisCredit(appcore, id, credit.IsFixed)
	appcore.Db.Delete(&credit)
	return credit
}

// CreateRecordsForCredit
// This function creates the records for a credit
// It is called from the controller
func CreateRecordsForCredit(appcore *webcore.AppCore, credit Credit) {
	// default values for the "credit" records
	id := strconv.Itoa(int(credit.ID))
	var comment string
	if credit.IsFixed {
		comment = "[FIXED] " + id
	} else {
		comment = "[CREDIT] " + id
	}
	amount_io := "out"
	base_date := carbon.Parse(credit.StartedAt)
	account_id := 3
	// si el gasto es fijo entonces los records son mutables.
	is_mutable := credit.IsFixed

	// Por cada uno de los payments, crear un record
	for i := 0; i < (credit.Payments); i++ {
		// primer dia del mes
		var record_date string
		if credit.IsFixed {
			record_date = base_date.AddMonths(i).StartOfDay().ToDateString()
		} else {
			record_date = base_date.AddMonths(i).StartOfMonth().StartOfDay().ToDateString()
		}
		CreateRecord(appcore, credit.Name, credit.Amount, amount_io, comment, record_date, credit.CategoryID, is_mutable, account_id)
	}
}

// DeleteAllRecordsOfThisCredit
// This function deletes all the records associated with a credit
// It is called from the controller
func deleteAllRecordsOfThisCredit(appcore *webcore.AppCore, credit_id int, is_fixed bool) {
	var records []Record
	if is_fixed {
		appcore.Db.Where("comment = '[FIXED] ?'", credit_id).Find(&records)
	} else {
		appcore.Db.Where("comment = '[CREDIT] ?'", credit_id).Find(&records)
	}
	appcore.Db.Delete(&records)
}

func getFinishAtDate(started_at string, add_months int) string {
	base_date := carbon.Parse(started_at)
	return base_date.AddMonths(add_months).StartOfMonth().StartOfDay().ToDateString()
}
