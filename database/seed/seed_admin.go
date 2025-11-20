package seed

import (
	"log"

	"github.com/rafaelchr/task-management/config"
	"github.com/rafaelchr/task-management/models"
	"github.com/rafaelchr/task-management/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name:     "Super Admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed to seed admin", err)
	} else {
		log.Println("Admin user seeded")
	}
}
