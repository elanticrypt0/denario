package models

import "github.com/elanticrypt0/denario_go/pkg/webcore"

// balances get the SUM of all records
// by month
// by year
// by month and year
// by category

// Balance struct
// type Balance struct {
// 	Year     int      `json:"year"`
// 	Month    int      `json:"month"`
// 	Records  []Record `json:"records"`
// 	Total    float64  `json:"total"`
// 	TotalIn  float64  `json:"total_in"`
// 	TotalOut float64  `json:"total_out"`
// }

// Find all balances
// Return the balances
// Return an error if there was a problem
// func FindAllBalances() []Balance {
// 	var balances []Balance
// 	appcore.Db.Order("year ASC, month ASC").Find(&balances)
// 	return balances
// }

type RecordBalance struct {
	Type  string  `json:"type"`
	Total float64 `json:"total"`
}

// Get total amount of a date by io
// Return the total amount
// Return an error if there was a problem
func GetTotalAmountOfDateByIo(appcore *webcore.AppCore, year, month, record_io string) RecordBalance {
	var balances []RecordBalance
	appcore.Db.Table("records").Select("SUM (amount) as total, amount_io as Type").Where("deleted_at is NULL and category_id NOT IN (28) and record_date LIKE ? AND amount_io = ?", year+"-"+month+"%", record_io).Group("amount_io").Find(&balances)

	if len(balances) == 0 {
		return RecordBalance{
			Type:  record_io,
			Total: 0,
		}
	} else {
		return balances[0]
	}

}

// Get total amount of a date by io
// Return the total amount
// Return an error if there was a problem
func GetTotalAmountOfDateByIoFixed(appcore *webcore.AppCore, year, month, record_io string) RecordBalance {
	var balances []RecordBalance
	appcore.Db.Table("records").Select("SUM (amount) as total, amount_io as Type").Where("deleted_at is NULL and category_id NOT IN (28) and record_date LIKE ? AND comment LIKE '[FIXED] %'", year+"-"+month+"%").Group("amount_io").Find(&balances)

	if len(balances) == 0 {
		return RecordBalance{
			Type:  record_io,
			Total: 0,
		}
	} else {
		return balances[0]
	}
}

// Get total amount of a date by io
// Return the total amount
// Return an error if there was a problem
func GetTotalAmountOfDateByIoCredits(appcore *webcore.AppCore, year, month, record_io string) RecordBalance {
	var balances []RecordBalance
	appcore.Db.Table("records").Select("SUM (amount) as total, amount_io as Type").Where("deleted_at is NULL and category_id NOT IN (28) and record_date LIKE ? AND comment LIKE '[CREDIT] %'", year+"-"+month+"%").Group("amount_io").Find(&balances)

	if len(balances) == 0 {
		return RecordBalance{
			Type:  record_io,
			Total: 0,
		}
	} else {
		return balances[0]
	}
}

// Get total amount of a date by io
// Return the total amount
// Return an error if there was a problem
func GetTotalAmountOfDateByIoSavings(appcore *webcore.AppCore, year, month string) RecordBalance {
	var balances []RecordBalance
	appcore.Db.Table("records").Select("SUM (amount) as total").Where("deleted_at is NULL and category_id NOT IN (28) and category_id=26 and amount_io='in' and record_date LIKE ?", year+"-"+month+"%").Group("amount_io").Find(&balances)

	if len(balances) == 0 {
		return RecordBalance{
			Type:  "in",
			Total: 0,
		}
	} else {
		return balances[0]
	}
}

// Get total amount of a date by io
// no suma si la categoria es una transferencia
// Return the total amount
// Return an error if there was a problem
func GetTotalAmountOfAccount(appcore *webcore.AppCore, year, month string, account_id int) RecordBalance {

	var account_in RecordBalance
	appcore.Db.Table("records").Select("SUM (amount) as total").Where("deleted_at is NULL and amount_io='in' and account_id =? and record_date LIKE ?", account_id, year+"-"+month+"%").Group("amount_io").Find(&account_in)

	var account_out RecordBalance
	appcore.Db.Table("records").Select("SUM (amount) as total").Where("deleted_at is NULL and amount_io='out' and account_id =? and record_date LIKE ?", account_id, year+"-"+month+"%").Group("amount_io").Find(&account_out)

	return RecordBalance{
		Type:  "diference",
		Total: account_in.Total - account_out.Total,
	}
}
