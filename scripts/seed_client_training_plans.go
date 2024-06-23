package scripts

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"ddx_hackathon_backend/models"

	"github.com/jinzhu/gorm"
)

func SeedClientTrainingPlans(db *gorm.DB) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var clients []models.User
	if err := db.Where("role = ?", "client").Find(&clients).Error; err != nil {
		log.Fatalf("Failed to fetch clients: %v", err)
	}

	for _, client := range clients {
		var clientTrainers []models.ClientTrainer
		if err := db.Where("client_id = ?", client.ID).Find(&clientTrainers).Error; err != nil {
			log.Fatalf("Failed to fetch client trainers for client %v: %v", client.ID, err)
		}

		trainerIDs := make([]uint, len(clientTrainers))
		for i, ct := range clientTrainers {
			trainerIDs[i] = ct.TrainerID
		}

		var trainingPlans []models.TrainingPlan
		if err := db.Where("created_by_id IN (?)", trainerIDs).Find(&trainingPlans).Error; err != nil {
			log.Fatalf("Failed to fetch training plans for trainers %v: %v", trainerIDs, err)
		}

		numPlans := rnd.Intn(len(trainingPlans) + 1)
		selectedPlans := rnd.Perm(len(trainingPlans))[:numPlans]

		for _, idx := range selectedPlans {
			plan := trainingPlans[idx]

			startDate := time.Now()
			endDate := time.Now().AddDate(0, 0, len(plan.Workouts)*2)

			clientPlan := models.ClientTrainingPlan{
				UserID:         client.ID,
				TrainingPlanID: plan.ID,
				Name:           plan.Name,
				Description:    plan.Description,
				StartDate:      &startDate,
				EndDate:        &endDate,
			}

			if err := db.Create(&clientPlan).Error; err != nil {
				log.Fatalf("Failed to create client training plan for client %v: %v", client.ID, err)
			}

			var workouts []models.Workout
			if err := db.Where("training_plan_id = ?", plan.ID).Order("id").Find(&workouts).Error; err != nil {
				log.Fatalf("Failed to fetch workouts for plan %v: %v", plan.ID, err)
			}

			for workoutOrder, workout := range workouts {
				startDate := time.Now().AddDate(0, 0, workoutOrder*2)
				endDate := time.Now().AddDate(0, 0, (workoutOrder*2)+1)

				clientWorkout := models.ClientWorkout{
					ClientTrainingPlanID: clientPlan.ID,
					WorkoutID:            workout.ID,
					Name:                 workout.Name,
					Description:          workout.Description,
					StartDate:            &startDate,
					EndDate:              &endDate,
					Order:                workoutOrder,
				}

				if err := db.Create(&clientWorkout).Error; err != nil {
					log.Fatalf("Failed to create client workout for client plan %v: %v", clientPlan.ID, err)
				}

				var workoutExercises []models.WorkoutExercise
				if err := db.Where("workout_id = ?", workout.ID).Order("id").Find(&workoutExercises).Error; err != nil {
					log.Fatalf("Failed to fetch workout exercises for workout %v: %v", workout.ID, err)
				}

				for exerciseOrder, workoutExercise := range workoutExercises {
					clientWorkoutExercise := models.ClientWorkoutExercise{
						ClientWorkoutID:   clientWorkout.ID,
						ExerciseID: 		   workoutExercise.ExerciseID,
						RestTime:          workoutExercise.RestTime,
						Order:             exerciseOrder,
					}

					if err := db.Create(&clientWorkoutExercise).Error; err != nil {
						log.Fatalf("Failed to create client workout exercise for client workout %v: %v", clientWorkout.ID, err)
					}

					var exerciseSets []models.ExerciseSet
					if err := db.Where("workout_exercise_id = ?", workoutExercise.ID).Order("id").Find(&exerciseSets).Error; err != nil {
						log.Fatalf("Failed to fetch exercise sets for workout exercise %v: %v", workoutExercise.ID, err)
					}

					for setOrder, set := range exerciseSets {
						clientExerciseSet := models.ClientExerciseSet{
							ClientWorkoutExerciseID: clientWorkoutExercise.ID,
							Reps:                    set.Reps,
							Duration:                set.Duration,
							RestTime:                set.RestTime,
							Order:                   setOrder,
						}

						if err := db.Create(&clientExerciseSet).Error; err != nil {
							log.Fatalf("Failed to create client exercise set for client workout exercise %v: %v", clientWorkoutExercise.ID, err)
						}
					}
				}
			}
		}
	}

	fmt.Println("Client training plans seeded successfully")
}
