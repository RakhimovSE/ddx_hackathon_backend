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

func TestGetTrainingPlans(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	// Create a training plan for testing
	db.Create(&models.TrainingPlan{Name: "Plan 1", Description: "Description 1"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/training_plans", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Plan 1")
}

func TestCreateTrainingPlan(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	trainingPlan := models.TrainingPlan{Name: "Plan 2", Description: "Description 2"}
	jsonValue, _ := json.Marshal(trainingPlan)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/training_plans", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var createdPlan models.TrainingPlan
	db.Where("name = ?", "Plan 2").First(&createdPlan)
	assert.Equal(t, "Plan 2", createdPlan.Name)
}

func TestDeleteTrainingPlan(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	// Create a training plan for testing
	plan := models.TrainingPlan{Name: "Plan 3", Description: "Description 3"}
	db.Create(&plan)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/training_plans/"+strconv.Itoa(int(plan.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var deletedPlan models.TrainingPlan
	err := db.Where("id = ?", plan.ID).First(&deletedPlan).Error
	assert.NotNil(t, err)
}

func TestUpdateTrainingPlan(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	// Create a training plan for testing
	plan := models.TrainingPlan{Name: "Plan 4", Description: "Description 4"}
	db.Create(&plan)

	updatedPlan := models.TrainingPlan{Name: "Updated Plan 4", Description: "Updated Description 4"}
	jsonValue, _ := json.Marshal(updatedPlan)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/training_plans/"+strconv.Itoa(int(plan.ID)), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var planFromDB models.TrainingPlan
	db.Where("id = ?", plan.ID).First(&planFromDB)
	assert.Equal(t, "Updated Plan 4", planFromDB.Name)
	assert.Equal(t, "Updated Description 4", planFromDB.Description)
}
