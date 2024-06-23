package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"ddx_hackathon_backend/models"
	"ddx_hackathon_backend/routes"
)

func TestGetClientExerciseSets(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := routes.SetupRouter(db)

	// Create test data
	client := models.User{Name: "Test Client", Email: "client@example.com", Password: "password"}
	trainer := models.User{Name: "Test Trainer", Email: "trainer@example.com", Password: "password"}
	db.Create(&client)
	db.Create(&trainer)

	clientTrainingPlan := models.ClientTrainingPlan{UserID: client.ID, Name: "Test Plan"}
	db.Create(&clientTrainingPlan)

	workout := models.Workout{Name: "Test Workout", TrainingPlanID: clientTrainingPlan.ID}
	db.Create(&workout)

	clientWorkout := models.ClientWorkout{ClientTrainingPlanID: clientTrainingPlan.ID, WorkoutID: workout.ID, Name: "Test Client Workout"}
	db.Create(&clientWorkout)

	exercise := models.Exercise{Name: "Test Exercise"}
	db.Create(&exercise)

	clientWorkoutExercise := models.ClientWorkoutExercise{ClientWorkoutID: clientWorkout.ID, ExerciseID: exercise.ID, RestTime: 60, Order: 1}
	db.Create(&clientWorkoutExercise)

	startDate1 := time.Now().Add(-1 * time.Hour)
	startDate2 := time.Now().Add(-2 * time.Hour)
	clientExerciseSet1 := models.ClientExerciseSet{ClientWorkoutExerciseID: clientWorkoutExercise.ID, Reps: 10, Duration: 60, RestTime: 30, Order: 1, StartDate: &startDate1}
	clientExerciseSet2 := models.ClientExerciseSet{ClientWorkoutExerciseID: clientWorkoutExercise.ID, Reps: 12, Duration: 50, RestTime: 30, Order: 2, StartDate: &startDate2}
	db.Create(&clientExerciseSet1)
	db.Create(&clientExerciseSet2)

	// Make the request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/clients/"+strconv.Itoa(int(client.ID))+"/exercise_sets/"+strconv.Itoa(int(exercise.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.ClientWorkoutExercise
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Check the response
	assert.Len(t, response, 1)
	assert.Len(t, response[0].Sets, 2)
	assert.Equal(t, clientExerciseSet1.ID, response[0].Sets[0].ID)
	assert.Equal(t, clientExerciseSet2.ID, response[0].Sets[1].ID)
}
