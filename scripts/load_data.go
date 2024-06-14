package scripts

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"ddx_hackathon_backend/models"

	"github.com/jinzhu/gorm"
)

type ExerciseData struct {
	Name             string   `json:"name"`
	Muscle           string   `json:"muscle"`
	AdditionalMuscle string   `json:"additionalMuscle"`
	Type             string   `json:"type"`
	Equipment        string   `json:"equipment"`
	Difficulty       string   `json:"difficulty"`
	Photos           []string `json:"photos"`
}

func LoadDataFromFile(db *gorm.DB) {
	file, err := os.Open("scripts/data/main_images.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var exercises []ExerciseData
	json.Unmarshal(byteValue, &exercises)

	for _, exercise := range exercises {
		// Handle Exercise Type
		var exType models.ExerciseType
		db.FirstOrCreate(&exType, models.ExerciseType{Name: exercise.Type})

		// Handle Difficulty
		var difficulty models.Difficulty
		db.FirstOrCreate(&difficulty, models.Difficulty{Level: exercise.Difficulty})

		// Create Exercise
		ex := models.Exercise{
			Name:         exercise.Name,
			Type:         exType,
			Difficulty:   difficulty,
		}
		db.Create(&ex)

		// Handle Muscles
		muscles := strings.Split(exercise.Muscle, ",")
		for _, muscleName := range muscles {
			var muscle models.Muscle
			db.FirstOrCreate(&muscle, models.Muscle{Name: strings.TrimSpace(muscleName)})
			db.Model(&ex).Association("Muscles").Append(&muscle)
		}

		// Handle Additional Muscles
		additionalMuscles := strings.Split(exercise.AdditionalMuscle, ",")
		for _, additionalMuscleName := range additionalMuscles {
			var additionalMuscle models.Muscle
			db.FirstOrCreate(&additionalMuscle, models.Muscle{Name: strings.TrimSpace(additionalMuscleName)})
			db.Model(&ex).Association("AdditionalMuscles").Append(&additionalMuscle)
		}

		// Handle Equipments
		equipments := strings.Split(exercise.Equipment, ",")
		for _, equipmentName := range equipments {
			var equipment models.Equipment
			db.FirstOrCreate(&equipment, models.Equipment{Name: strings.TrimSpace(equipmentName)})
			db.Model(&ex).Association("Equipments").Append(&equipment)
		}

		// Handle Photos
		for _, photo := range exercise.Photos {
			exercisePhoto := models.ExercisePhoto{
				ExerciseID: ex.ID,
				PhotoURL:   photo,
			}
			db.Create(&exercisePhoto)
		}
	}

	fmt.Println("Data loaded successfully")
}
