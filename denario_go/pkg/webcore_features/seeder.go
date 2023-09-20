package webcore_features

import (
	"github.com/gofiber/fiber/v2"
)

func Seed(c *fiber.Ctx) error {
	// seedCategories()
	// seedCredits()
	// seedRecords()
	return c.JSON("OK")
}

func seedTable(table_name string) {

}

// func seedCategories() {
// 	file := helpers.ReadJsonFile("categories")
// 	cat := []models.Category{}
// 	err := json.Unmarshal(file, &cat)
// 	if err != nil {
// 		log.Fatal("Cant unmarshal json", err)
// 	}
// 	for _, category := range cat {
// 		models.CreateCategory(category.Name)
// 	}
// 	log.Println("Categories seeded")
// }

// func seedCredits() {
// 	file := helpers.ReadJsonFile("credits")
// 	toInsert := []models.Credit{}
// 	err := json.Unmarshal(file, &toInsert)
// 	if err != nil {
// 		log.Fatal("Cant unmarshal json", err)
// 	}
// 	for _, item := range toInsert {
// 		models.CreateCredit(item.Name, item.Comment, item.Amount, item.Payments, item.StartedAt, item.CategoryID, item.IsFixed)
// 	}
// 	log.Println("Credits seeded")
// }

// func seedRecords() {
// 	file := helpers.ReadJsonFile("records")
// 	toInsert := []models.Record{}
// 	err := json.Unmarshal(file, &toInsert)
// 	if err != nil {
// 		log.Fatal("Cant unmarshal json", err)
// 	}
// 	for _, item := range toInsert {
// 		models.CreateRecord(item.Name, item.Amount, item.AmountIo, item.Comment, item.RecordDate, item.CategoryID, item.IsMutable, item.AccountID)
// 	}
// 	log.Println("Records seeded")
// }
