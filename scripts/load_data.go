package scripts

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

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
		ex := models.Exercise{
			Name:             exercise.Name,
			Muscle:           models.Muscle(exercise.Muscle),
			AdditionalMuscle: models.Muscle(exercise.AdditionalMuscle),
			Type:             models.ExerciseType(exercise.Type),
			Equipment:        models.Equipment(exercise.Equipment),
			Difficulty:       models.Difficulty(exercise.Difficulty),
		}

		db.Create(&ex)

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
