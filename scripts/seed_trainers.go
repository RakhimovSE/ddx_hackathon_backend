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

var trainerFirstNames = []string{
	"Иван", "Мария", "Анна", "Дмитрий", "Екатерина",
	"Сергей", "Елена", "Алексей", "Ольга", "Владимир",
	"Павел", "Юлия", "Виктор", "Ирина", "Андрей",
	"Татьяна", "Роман", "Лариса", "Михаил", "Светлана",
	"Николай", "Алина", "Артур", "Евгения", "Григорий",
	"Дарья", "Максим", "Людмила", "Валентин", "Наталья",
}

var trainerLastNames = []string{
	"Иванов", "Петрова", "Смирнова", "Кузнецов", "Волкова",
	"Морозов", "Соколова", "Попов", "Лебедева", "Козлов",
	"Новиков", "Федорова", "Михайлов", "Борисова", "Тихонов",
	"Гусева", "Крылов", "Захарова", "Ершов", "Макарова",
	"Никитин", "Жукова", "Шмидт", "Комарова", "Баранов",
	"Белов", "Фролова", "Капустин", "Демидова", "Гаврилов",
}

var trainerBios = []string{
	"Люблю активный образ жизни и помогаю другим достигать их целей.",
	"Тренировки – это моя страсть, я делаю их интересными и эффективными.",
	"Я помогу вам стать лучше с каждым днем.",
	"Моя цель – сделать фитнес доступным и приятным для всех.",
	"Я верю, что каждый может достичь успеха в тренировках.",
	"Здоровье и фитнес – моя профессия и образ жизни.",
	"Я покажу вам, как добиться результата с удовольствием.",
	"Вдохновляю своих клиентов на постоянные тренировки и самосовершенствование.",
	"Стремлюсь сделать каждую тренировку интересной и продуктивной.",
	"Обожаю спорт и мотивирую других к здоровому образу жизни.",
	"Моя работа - это моя страсть, каждый день я работаю над собой и помогаю другим.",
	"Профессиональный подход и индивидуальные программы тренировок для каждого.",
	"Опытный тренер, готовый помочь в достижении ваших целей.",
	"Работаю с разными уровнями подготовки, от новичков до профессионалов.",
	"Моя цель - сделать так, чтобы вы полюбили спорт так же, как и я.",
	"Каждая тренировка - это шаг к лучшей версии себя.",
	"Постоянно обучаюсь и развиваюсь, чтобы быть лучшим тренером для своих клиентов.",
	"Создаю индивидуальные программы, которые учитывают все особенности клиента.",
	"Верю, что каждый может достичь своих целей, нужно только начать.",
	"Сделаю все, чтобы тренировки приносили радость и были максимально эффективными.",
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

		email := faker.Email()
		avatarUrl := helpers.GenerateAvatarURL(email, "trainer")
		trainer := models.User{
			Name:      randomName(rnd, trainerFirstNames, trainerLastNames),
			Email:     email,
			Password:  string(hashedPassword),
			Role:      "trainer",
			AvatarUrl: &avatarUrl,
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
