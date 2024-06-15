package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"

	"ddx_hackathon_backend/models"
	"ddx_hackathon_backend/routes"
)

func TestGetExercises(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	exerciseType := models.ExerciseType{Name: "Type 1"}
	difficulty := models.Difficulty{Level: "Difficulty 1"}
	db.Create(&exerciseType)
	db.Create(&difficulty)

	exercise1 := models.Exercise{Name: "Exercise 1", TypeID: exerciseType.ID, DifficultyID: difficulty.ID, Type: exerciseType, Difficulty: difficulty}
	exercise2 := models.Exercise{Name: "Exercise 2", TypeID: exerciseType.ID, DifficultyID: difficulty.ID, Type: exerciseType, Difficulty: difficulty}
	db.Create(&exercise1)
	db.Create(&exercise2)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exercises", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Exercise 1")
	assert.Contains(t, w.Body.String(), "Exercise 2")
}

func TestCreateExercise(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	exerciseType := models.ExerciseType{Name: "Type 1"}
	difficulty := models.Difficulty{Level: "Difficulty 1"}
	db.Create(&exerciseType)
	db.Create(&difficulty)

	muscle := models.Muscle{Name: "Muscle 1"}
	additionalMuscle := models.Muscle{Name: "Additional Muscle 1"}
	equipment := models.Equipment{Name: "Equipment 1"}
	db.Create(&muscle)
	db.Create(&additionalMuscle)
	db.Create(&equipment)

	exercise := models.Exercise{
		Name:              "New Exercise",
		TypeID:            exerciseType.ID,
		DifficultyID:      difficulty.ID,
		Muscles:           []models.Muscle{muscle},
		AdditionalMuscles: []models.Muscle{additionalMuscle},
		Equipments:        []models.Equipment{equipment},
	}
	jsonValue, _ := json.Marshal(exercise)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/exercises", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var createdExercise models.Exercise
	db.Preload("Muscles").Preload("AdditionalMuscles").Preload("Equipments").Where("name = ?", "New Exercise").First(&createdExercise)
	assert.Equal(t, "New Exercise", createdExercise.Name)
	assert.Equal(t, 1, len(createdExercise.Muscles))
	assert.Equal(t, 1, len(createdExercise.AdditionalMuscles))
	assert.Equal(t, 1, len(createdExercise.Equipments))
}

func TestDeleteExercise(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	exerciseType := models.ExerciseType{Name: "Type 1"}
	difficulty := models.Difficulty{Level: "Difficulty 1"}
	db.Create(&exerciseType)
	db.Create(&difficulty)

	exercise := models.Exercise{Name: "Exercise to Delete", TypeID: exerciseType.ID, DifficultyID: difficulty.ID, Type: exerciseType, Difficulty: difficulty}
	db.Create(&exercise)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/exercises/"+strconv.Itoa(int(exercise.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var deletedExercise models.Exercise
	err := db.Where("id = ?", exercise.ID).First(&deletedExercise).Error
	assert.NotNil(t, err)
}

func TestUpdateExercise(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	exerciseType := models.ExerciseType{Name: "Type 1"}
	difficulty := models.Difficulty{Level: "Difficulty 1"}
	db.Create(&exerciseType)
	db.Create(&difficulty)

	muscle := models.Muscle{Name: "Muscle 1"}
	additionalMuscle := models.Muscle{Name: "Additional Muscle 1"}
	equipment := models.Equipment{Name: "Equipment 1"}
	db.Create(&muscle)
	db.Create(&additionalMuscle)
	db.Create(&equipment)

	exercise := models.Exercise{
		Name:              "Exercise to Update",
		TypeID:            exerciseType.ID,
		DifficultyID:      difficulty.ID,
		Muscles:           []models.Muscle{muscle},
		AdditionalMuscles: []models.Muscle{additionalMuscle},
		Equipments:        []models.Equipment{equipment},
	}
	db.Create(&exercise)

	updatedExercise := models.Exercise{
		Name:              "Updated Exercise",
		TypeID:            exerciseType.ID,
		DifficultyID:      difficulty.ID,
		Muscles:           []models.Muscle{muscle},
		AdditionalMuscles: []models.Muscle{additionalMuscle},
		Equipments:        []models.Equipment{equipment},
	}
	jsonValue, _ := json.Marshal(updatedExercise)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/exercises/"+strconv.Itoa(int(exercise.ID)), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var exerciseFromDB models.Exercise
	db.Preload("Muscles").Preload("AdditionalMuscles").Preload("Equipments").Where("id = ?", exercise.ID).First(&exerciseFromDB)
	assert.Equal(t, "Updated Exercise", exerciseFromDB.Name)
}
