package initializers

import (
	"log"
	"os"

	"github.com/azam-sh/books/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}

func Seed() {
	roles := []models.Role{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "user"},
	}
	for _, r := range roles {
		DB.Create(&r)
	}

	accesses := []models.Access{
		{ID: 1, Name: "free"},
		{ID: 2, Name: "silver"},
		{ID: 3, Name: "gold"},
	}
	for _, a := range accesses {
		DB.Create(&a)
	}

	categories := []models.Category{
		{ID: 1, Name: "romance"},
		{ID: 2, Name: "mysteries"},
		{ID: 3, Name: "history"},
		{ID: 4, Name: "business"},
		{ID: 5, Name: "biographies"},
	}
	for _, c := range categories {
		DB.Create(&c)
	}
}

func SyncDatabase() {
	err := DB.AutoMigrate(&models.Role{}, &models.Access{}, &models.User{}, &models.Book{})
	if err != nil {
		log.Fatal("migration failed")
	}
	Seed()
}
