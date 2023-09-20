package models

import (
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Name       string   `json:"name"`
	Amount     float64  `json:"amount"`
	AmountIo   string   `json:"amount_io"`
	Comment    string   `json:"comment"`
	RecordDate string   `json:"record_date"`
	CategoryID int      `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID"`
	IsMutable  bool     `json:"is_mutable"`
	AccountID  int      `json:"account_id"`
	Account    Account  `gorm:"foreignKey:AccountID"`
}

// Find a record by ID
// Return the record
// Return an error if there was a problem
func FindOneRecord(appcore *webcore.AppCore, id int) Record {
	var record Record
	appcore.Db.Preload("Category").Preload("Account").First(&record, id)
	return record
}

// Find all records
// Return the records
// Return an error if there was a problem
func FindAllRecords(appcore *webcore.AppCore) []Record {
	var records []Record
	// appcore.Db.Order("record_date ASC").Find(&records)
	appcore.Db.Preload("Category").Preload("Account").Order("record_date ASC").Find(&records)
	return records
}

// Create a new record
// Return the new record
// Return an error if there was a problem
func CreateRecord(appcore *webcore.AppCore, name string, amount float64, amount_io string, comment string, record_date string, category_id int, is_mutable bool, account_id int) Record {
	record := Record{
		Name:       name,
		Amount:     amount,
		AmountIo:   amount_io,
		Comment:    comment,
		RecordDate: record_date,
		CategoryID: category_id,
		IsMutable:  is_mutable,
		AccountID:  account_id,
	}
	appcore.Db.Create(&record)
	return record
}

// Update a record if it is mutable
// Return the updated record
// Return an error if there was a problem
// TODO funciona pero devuelve el record actualizado. Deberìa mostrar el record viejo
func UpdateRecord(appcore *webcore.AppCore, new_record Record) Record {
	current_record := new(Record)
	appcore.Db.First(&current_record, new_record.ID)
	if current_record.IsMutable == false {
		return *current_record
	}
	appcore.Db.Save(&new_record)
	return new_record
}

// Delete a record
// Return the deleted record
// Return an error if there was a problem
func DeleteRecord(appcore *webcore.AppCore, id int) Record {
	record := Record{}
	appcore.Db.First(&record, id)
	appcore.Db.Delete(&record)
	return record
}

// find by year and month

// Find all records by year and month
// Return the records
// Return an error if there was a problem
func FindRecordsByYearAndMonth(appcore *webcore.AppCore, year, month, record_io string) []Record {
	var records []Record

	if record_io == "" {
		appcore.Db.Preload("Category").Preload("Account").Where("record_date LIKE ?", year+"-"+month+"%").Order("record_date ASC").Find(&records)
	} else {
		appcore.Db.Preload("Category").Preload("Account").Where("record_date LIKE ? AND amount_io = ?", year+"-"+month+"%", record_io).Order("record_date ASC").Find(&records)
	}

	return records
}
