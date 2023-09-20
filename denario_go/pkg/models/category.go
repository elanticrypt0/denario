package models

import (
	"github.com/elanticrypt0/denario_go/pkg/webcore"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func FindOneCategory(appcore *webcore.AppCore, id int) Category {
	var category Category
	appcore.Db.First(&category, id)
	return category
}

func FindAllCategories(appcore *webcore.AppCore) []Category {
	var categories []Category
	// TODO revisar esta parte
	// appcore.Db.Order("created_at ASC").Find(&categories)
	appcore.Db.Order("created_at ASC").Find(&categories)
	return categories
}

func CreateCategory(appcore *webcore.AppCore, name string) Category {
	category := Category{
		Name: name,
	}
	appcore.Db.Create(&category)
	return category
}

func UpdateCategory(appcore *webcore.AppCore, category Category) Category {
	appcore.Db.Save(&category)
	return category
}

func DeleteCategory(appcore *webcore.AppCore, id int) Category {
	// var category Category
	category := Category{}
	appcore.Db.First(&category, id)
	appcore.Db.Delete(&category)
	return category
}
