package scripts

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"ddx_hackathon_backend/helpers"
	"ddx_hackathon_backend/models"

	"github.com/go-faker/faker/v4"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var clientFirstNames = []string{
	"Алессандро", "Лука", "Федерико", "Маттео", "Джулия",
	"Франческа", "Марко", "Леонардо", "Элиза", "Роберто",
	"Антонио", "Джованни", "Франческо", "Марио", "Валентина",
	"София", "Джузеппе", "Сильвия", "Винченцо", "Никола",
	"Андреа", "Фабио", "Стефания", "Энрико", "Габриэла",
	"Давиде", "Ренато", "Клаудио", "Моника", "Лучия",
}

var clientLastNames = []string{
	"Росси", "Феррари", "Бьянки", "Риччи", "Морети",
	"Костантини", "Капелли", "Пальмиери", "Джентиле", "Лонго",
	"Марини", "Коста", "Фонтана", "Джакоббе", "Ломбарди",
	"Барбьери", "Каппуччи", "Сантини", "Ринальди", "Донати",
	"Сартори", "Марчезе", "Пеллегрини", "Негри", "Ферретти",
	"Де Лука", "Орландо", "Скалья", "Джордани", "Карузо",
}

func SeedClients(db *gorm.DB) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := "1"

	for i := 0; i < 50; i++ {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		email := faker.Email()
		avatarUrl := helpers.GenerateAvatarURL(email, "client")
		client := models.User{
			Name:      randomName(rnd, clientFirstNames, clientLastNames),
			Email:     email,
			Password:  string(hashedPassword),
			Role:      "client",
			AvatarUrl: &avatarUrl,
		}

		if err := db.Create(&client).Error; err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
	}

	fmt.Println("20 clients added successfully")
}
