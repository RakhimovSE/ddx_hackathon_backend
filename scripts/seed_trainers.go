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

var trainerFirstNames = []string{
	"Иван", "Мария", "Анна", "Дмитрий", "Екатерина",
	"Сергей", "Елена", "Алексей", "Ольга", "Владимир",
}

var trainerLastNames = []string{
	"Иванов", "Петрова", "Смирнова", "Кузнецов", "Волкова",
	"Морозов", "Соколова", "Попов", "Лебедева", "Козлов",
}

var trainerBios = []string{
	"Люблю активный образ жизни и помогаю другим достигать их целей.",
	"Тренировки – это моя страсть, я делаю их интересными и эффективными.",
	"Я помогу вам стать лучше с каждым днем.",
	"Моя цель – сделать фитнес доступным и приятным для всех.",
	"Я верю, что каждый может достичь успеха в тренировках.",
	"Здоровье и фитнес – моя профессия и образ жизни.",
	"Я покажу вам, как добиться результата с удовольствием.",
}

func randomName(r *rand.Rand, firstNames []string, lastNames []string) string {
	return randomElement(r, firstNames) + " " + randomElement(r, lastNames)
}

func randomSpecialties(db *gorm.DB, specialties []string, n int, rnd *rand.Rand) []models.Specialty {
	rnd.Shuffle(len(specialties), func(i, j int) { 
		specialties[i], specialties[j] = specialties[j], specialties[i] 
	})
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
			Name:      randomName(rnd, trainerFirstNames, trainerLastNames),
			Email:     faker.Email(),
			Password:  string(hashedPassword),
			Role:      "trainer",
			TrainerProfile: &models.TrainerProfile{
				Specialties: randomSpecialties(db, specialties, rnd.Intn(3)+1, rnd),
				Experience:  rnd.Intn(120), // Random experience between 0 and 120 months
				Bio:         randomElement(rnd, trainerBios),
			},
		}

		if err := db.Create(&trainer).Error; err != nil {
			log.Fatalf("Failed to create trainer: %v", err)
		}
	}

	fmt.Println("20 trainers added successfully")
}
