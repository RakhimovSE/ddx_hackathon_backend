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

type ExerciseExtensionData struct {
	Unit       string `json:"unit"`
	NeedWeight bool   `json:"needWeight"`
}

type ExerciseExtensionMap map[string]ExerciseExtensionData

func LoadDataFromFile(db *gorm.DB) {
	// Load main exercise data
	mainFile, err := os.Open("scripts/data/main_images.json")
	if err != nil {
		log.Fatalf("Failed to open main file: %v", err)
	}
	defer mainFile.Close()

	mainByteValue, _ := io.ReadAll(mainFile)

	var exercises []ExerciseData
	json.Unmarshal(mainByteValue, &exercises)

	// Load exercise extension data
	extensionFile, err := os.Open("scripts/data/exercises_extensions.json")
	if err != nil {
		log.Fatalf("Failed to open extensions file: %v", err)
	}
	defer extensionFile.Close()

	extensionByteValue, _ := io.ReadAll(extensionFile)

	var exerciseExtensions ExerciseExtensionMap
	json.Unmarshal(extensionByteValue, &exerciseExtensions)

	for _, exercise := range exercises {
		// Handle Exercise Type
		var exType models.ExerciseType
		db.FirstOrCreate(&exType, models.ExerciseType{Name: exercise.Type})

		// Handle Difficulty
		var difficulty models.Difficulty
		db.FirstOrCreate(&difficulty, models.Difficulty{Level: exercise.Difficulty})

		// Determine Unit and NeedWeight
		extensionData, exists := exerciseExtensions[exercise.Name]
		unit := "reps"        // Default unit
		needWeight := false   // Default needWeight
		if exists {
			unit = extensionData.Unit
			needWeight = extensionData.NeedWeight
		}

		// Create Exercise
		ex := models.Exercise{
			Name:        exercise.Name,
			Type:        exType,
			Difficulty:  difficulty,
			Unit:        unit,
			NeedWeight:  needWeight,
			SourceType:  "catalog",
			CreatedByID: nil, // Since these exercises are from catalog, they are not user-created
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
			if strings.TrimSpace(equipmentName) == "Отсутствует" {
				equipmentName = "Без оборудования"
			}
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
