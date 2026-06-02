package database

import (
	"edart-core/app/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	DB = db
	testQuery()
}

func testQuery() {

	var rows []models.TestTable

	result := DB.Find(&rows)
	// DB.AutoMigrate(
	// 	&models.TxnSuccess{},
		// &models.DailyLedger{},
		// &models.AssetListing{},
		// &models.AccountBalance{},
	// 	&models.UserInfo{},
	// 	&models.BalanceLedger{},
	// 	&models.TestTable{},
	// )

	if result.Error != nil {
		log.Println("Query failed:", result.Error)
		return
	}

	log.Println("Query success. Rows found:", len(rows))

	for _, r := range rows {
		log.Printf("ID: %d | Text: %s\n", r.ID, r.Text)
	}
}