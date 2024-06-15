package scripts

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"ddx_hackathon_backend/models"

	"github.com/go-faker/faker/v4"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

func randomSpecialties(db *gorm.DB, specialties []string, n int, rnd *rand.Rand) []models.Specialty {
	rnd.Shuffle(len(specialties), func(i, j int) { specialties[i], specialties[j] = specialties[j], specialties[i] })
	result := make([]models.Specialty, n)
	for i := 0; i < n; i++ {
		var specialty models.Specialty
		db.FirstOrCreate(&specialty, models.Specialty{Name: specialties[i]})
		result[i] = specialty
	}
	return result
}

func SeedTrainers(db *gorm.DB) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	specialties := []string{"Йога", "Похудение", "Скалолазание", "Фитнес", "Кроссфит", "Пилатес", "Бокс", "Стретчинг", "Аэробика", "Танцы"}
	password := "1"

	for i := 0; i < 20; i++ {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		trainer := models.User{
			Name:      faker.Name(),
			Email:     faker.Email(),
			Password:  string(hashedPassword),
			Role:      "trainer",
			TrainerProfile: &models.TrainerProfile{
				Specialties: randomSpecialties(db, specialties, rnd.Intn(3)+1, rnd),
				Experience:  rnd.Intn(120), // Random experience between 0 and 120 months
				Bio:         faker.Sentence(),
			},
		}

		if err := db.Create(&trainer).Error; err != nil {
			log.Fatalf("Failed to create trainer: %v", err)
		}
	}

	fmt.Println("20 trainers added successfully")
}
