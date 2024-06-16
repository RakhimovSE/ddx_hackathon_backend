package tests

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"

	"ddx_hackathon_backend/models"
	"ddx_hackathon_backend/routes"
)

func TestGetClientWorkouts(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	client := models.User{Name: "Test Client", Email: "client@example.com", Password: "password", Role: "client"}
	db.Create(&client)

	trainingPlan := models.TrainingPlan{Name: "Test Plan", Description: "Test Description", CreatedByID: &client.ID}
	db.Create(&trainingPlan)

	workout := models.Workout{TrainingPlanID: trainingPlan.ID, Name: "Workout 1", Description: "Description 1"}
	db.Create(&workout)

	clientTrainingPlan := models.ClientTrainingPlan{
		UserID:         client.ID,
		TrainingPlanID: trainingPlan.ID,
		Name:           "Test Plan",
		Description:    "Test Description",
		StartDate:      timePtr(time.Now()),
		EndDate:        timePtr(time.Now().AddDate(0, 1, 0)),
	}
	db.Create(&clientTrainingPlan)

	clientWorkout1 := models.ClientWorkout{
		ClientTrainingPlanID: clientTrainingPlan.ID,
		WorkoutID:            workout.ID,
		Name:                 "Workout 1",
		Description:          "Description 1",
		Order:                1,
		StartDate:            timePtr(time.Now().AddDate(0, 0, 1)),
		EndDate:              timePtr(time.Now().AddDate(0, 0, 1).Add(time.Hour)),
	}
	db.Create(&clientWorkout1)

	clientWorkout2 := models.ClientWorkout{
		ClientTrainingPlanID: clientTrainingPlan.ID,
		WorkoutID:            workout.ID,
		Name:                 "Workout 2",
		Description:          "Description 2",
		Order:                2,
		StartDate:            timePtr(time.Now().AddDate(0, 0, -1)),
		EndDate:              timePtr(time.Now().AddDate(0, 0, -1).Add(time.Hour)),
	}
	db.Create(&clientWorkout2)

	// Test fetching all workouts
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/clients/"+strconv.Itoa(int(client.ID))+"/workouts", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Workout 1")
	assert.Contains(t, w.Body.String(), "Workout 2")

	// Test fetching upcoming workouts
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/clients/"+strconv.Itoa(int(client.ID))+"/workouts?status=upcoming", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Workout 1")
	assert.NotContains(t, w.Body.String(), "Workout 2")

	// Test fetching past workouts
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/clients/"+strconv.Itoa(int(client.ID))+"/workouts?status=past", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Workout 2")
	assert.NotContains(t, w.Body.String(), "Workout 1")

	// Test fetching workouts by date
	w = httptest.NewRecorder()
	dateStr := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	req, _ = http.NewRequest("GET", "/clients/"+strconv.Itoa(int(client.ID))+"/workouts?date="+dateStr, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Workout 1")
	assert.NotContains(t, w.Body.String(), "Workout 2")
}

func timePtr(t time.Time) *time.Time {
	return &t
}
