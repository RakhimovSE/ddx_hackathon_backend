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
}

var planDescriptions = []string{
	"Этот план разработан для увеличения силы и выносливости.",
	"Идеально подходит для тех, кто только начинает свои тренировки.",
	"Йога для всех уровней подготовки.",
	"Кардио тренировка для сжигания калорий.",
	"План для улучшения гибкости и растяжки.",
}

var workoutDescriptions = []string{
	"Отличная тренировка для начала дня.",
	"Идеально для улучшения вашей силы.",
	"Проведите время с пользой для тела и души.",
	"Энергичная тренировка для сжигания жира.",
	"Улучшите свою гибкость и координацию.",
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
		numPlans := rnd.Intn(4) // Up to 4 training plans per trainer
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
					Name:           fmt.Sprintf("Тренировка #%d", k+1),
					Description:    randomElement(rnd, workoutDescriptions),
					DaysUntilNext:  rnd.Intn(4) + 1, // 1-3 days until next workout
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
					}

					if err := db.Create(&workoutExercise).Error; err != nil {
						log.Fatalf("Failed to create workout exercise: %v", err)
					}

					numSets := rnd.Intn(3) + 3 // 3-5 sets per exercise
					for m := 0; m < numSets; m++ {
						set := models.ExerciseSet{
							WorkoutExerciseID: workoutExercise.ID,
							RestTime:          30, // Default rest time between sets
						}

						if rnd.Intn(2) == 0 {
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
