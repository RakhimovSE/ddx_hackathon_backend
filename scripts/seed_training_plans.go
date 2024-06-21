package scripts

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"ddx_hackathon_backend/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var planNames = []string{
	"Сила и выносливость",
	"Тренировка для начинающих",
	"Йога для всех",
	"Кардио тренировка",
	"Тренировка на гибкость",
	"Функциональный тренинг",
	"Тренировка на выносливость",
	"Силовая тренировка",
	"Интервальная тренировка",
	"Тренировка на силу",
	"Тренировка с весом",
	"Тренировка для похудения",
	"Тренировка на растяжку",
	"Здоровая спина",
	"Здоровые суставы",
	"Общая физическая подготовка",
	"Йога для начинающих",
	"Пилатес",
	"Танцевальная тренировка",
	"Кроссфит тренировка",
	"Круговая тренировка",
	"Раннее утро",
	"Легкая тренировка",
	"Быстрая тренировка",
	"Тренировка на свежем воздухе",
	"Домашняя тренировка",
	"Силовые нагрузки",
	"Тренировка на пресс",
	"Зарядка",
	"Вечерняя тренировка",
}

var planDescriptions = []string{
	"Этот план разработан для увеличения силы и выносливости.",
	"Идеально подходит для тех, кто только начинает свои тренировки.",
	"Йога для всех уровней подготовки.",
	"Кардио тренировка для сжигания калорий.",
	"План для улучшения гибкости и растяжки.",
	"Функциональные тренировки для всех.",
	"Увеличьте свою выносливость с помощью этого плана.",
	"Тренировки для увеличения силы.",
	"Интервальные тренировки для максимального эффекта.",
	"Развитие силы с помощью специализированных упражнений.",
	"Использование собственного веса для тренировки.",
	"Идеальный план для снижения веса.",
	"Улучшите свою растяжку и гибкость.",
	"Поддержите здоровье своей спины.",
	"Укрепление суставов с помощью упражнений.",
	"Общая подготовка тела для всех.",
	"Йога для начинающих всех возрастов.",
	"Пилатес для укрепления тела.",
	"Танцевальные движения для тренировки.",
	"Кроссфит для профессионалов.",
}

var workoutDescriptions = []string{
	"Отличная тренировка для начала дня.",
	"Идеально для улучшения вашей силы.",
	"Проведите время с пользой для тела и души.",
	"Энергичная тренировка для сжигания жира.",
	"Улучшите свою гибкость и координацию.",
	"Тренировка, которая поможет вам проснуться.",
	"Тренировка, направленная на укрепление мышц.",
	"Занимайтесь где угодно и когда угодно.",
	"Тренировка для всего тела.",
	"Тренировка, которую можно выполнять дома.",
	"Быстрая тренировка для занятых людей.",
	"Тренировка, которая поддержит ваше здоровье.",
	"Легкая тренировка для восстановления.",
	"Эффективные упражнения для вашего пресса.",
	"Утренняя зарядка для бодрости.",
	"Вечерняя тренировка для расслабления.",
	"Тренировка на свежем воздухе.",
	"Домашняя тренировка без оборудования.",
	"Тренировка для всех уровней подготовки.",
	"Интенсивная тренировка для профессионалов.",
	"Тренировка с весами.",
	"Силовая тренировка для начинающих.",
	"Функциональная тренировка для повседневной жизни.",
	"Тренировка на выносливость и силу.",
	"Быстрая зарядка перед началом дня.",
	"Вечерняя йога для расслабления.",
	"Кардио тренировка для снижения веса.",
	"Тренировка для увеличения гибкости.",
	"Общая тренировка для всего тела.",
	"Тренировка на свежем воздухе для бодрости.",
}

func SeedTrainingPlans(db *gorm.DB) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var trainers []models.User
	if err := db.Where("role = ?", "trainer").Find(&trainers).Error; err != nil {
		log.Fatalf("Failed to fetch trainers: %v", err)
	}

	var exercises []models.Exercise
	if err := db.Find(&exercises).Error; err != nil {
		log.Fatalf("Failed to fetch exercises: %v", err)
	}

	for _, trainer := range trainers {
		numPlans := rnd.Intn(4) + 1 // 1-4 training plans per trainer
		for j := 0; j < numPlans; j++ {
			plan := models.TrainingPlan{
				Name:        randomElement(rnd, planNames),
				Description: randomElement(rnd, planDescriptions),
				CreatedByID: &trainer.ID,
			}

			if err := db.Create(&plan).Error; err != nil {
				log.Fatalf("Failed to create training plan: %v", err)
			}

			numWorkouts := rnd.Intn(4) + 12 // 12-15 workouts per plan
			for k := 0; k < numWorkouts; k++ {
				workout := models.Workout{
					TrainingPlanID: plan.ID,
					Name:           randomElement(rnd, planNames),
					Description:    randomElement(rnd, workoutDescriptions),
					DaysUntilNext:  rnd.Intn(4) + 1, // 1-3 days until next workout
					Order:          k + 1, // Порядок тренировки в плане
				}

				if err := db.Create(&workout).Error; err != nil {
					log.Fatalf("Failed to create workout: %v", err)
				}

				numExercises := rnd.Intn(6) + 5 // 5-10 exercises per workout
				for l := 0; l < numExercises; l++ {
					exercise := exercises[rnd.Intn(len(exercises))]
					workoutExercise := models.WorkoutExercise{
						WorkoutID:  workout.ID,
						ExerciseID: exercise.ID,
						RestTime:   60, // Default rest time between exercises
						Order:      l + 1, // Порядок упражнения в тренировке
					}

					if err := db.Create(&workoutExercise).Error; err != nil {
						log.Fatalf("Failed to create workout exercise: %v", err)
					}

					numSets := rnd.Intn(3) + 3 // 3-5 sets per exercise
					for m := 0; m < numSets; m++ {
						set := models.ExerciseSet{
							WorkoutExerciseID: workoutExercise.ID,
							RestTime:          30, // Default rest time between sets
							Order:             m + 1, // Порядок подхода в упражнении
						}

						if exercise.Unit == "reps" {
							set.Reps = rnd.Intn(15) + 5 // 5-20 reps
						} else {
							set.Duration = rnd.Intn(45) + 15 // 15-60 seconds
						}

						if err := db.Create(&set).Error; err != nil {
							log.Fatalf("Failed to create exercise set: %v", err)
						}
					}
				}
			}
		}
	}

	fmt.Println("Training plans seeded successfully")
}
