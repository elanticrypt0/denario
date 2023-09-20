package models

import (
	"strings"

	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name   string `json:"name"`
	Short  string `json:"short"`
	Detail string `json:"detail"`
}

func FindOneAccount(appcore *webcore.AppCore, id int) Account {
	var account Account
	appcore.Db.First(&account, id)
	return account
}

func FindAllAccounts(appcore *webcore.AppCore) []Account {
	var accounts []Account
	// TODO revisar esta parte
	appcore.Db.Order("created_at ASC").Find(&accounts)
	return accounts
}

func CreateAccount(appcore *webcore.AppCore, name, short, detail string) Account {
	account := Account{
		Name:   name,
		Short:  strings.ToUpper(short),
		Detail: detail,
	}
	appcore.Db.Create(&account)
	return account
}

func UpdateAccount(appcore *webcore.AppCore, account Account) Account {
	appcore.Db.Save(&account)
	return account
}

func DeleteAccount(appcore *webcore.AppCore, id int) Account {
	var account Account
	appcore.Db.First(&account, id)
	appcore.Db.Delete(&account)
	return account
}
