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
	"golang.org/x/crypto/bcrypt"

	"ddx_hackathon_backend/models"
)

func TestGetUsers(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := SetupRouter(db)

	// Create a user for testing
	db.Create(&models.User{Name: "John Doe", Email: "john.doe@example.com", Password: "password"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestCreateUser(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := SetupRouter(db)

	user := models.User{Name: "Jane Doe", Email: "jane.doe@example.com", Password: "password"}
	jsonValue, _ := json.Marshal(user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var createdUser models.User
	db.Where("email = ?", "jane.doe@example.com").First(&createdUser)
	assert.Equal(t, "Jane Doe", createdUser.Name)
}

func TestDeleteUser(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := SetupRouter(db)

	// Create a user for testing
	user := models.User{Name: "John Doe", Email: "john.doe@example.com", Password: "password"}
	db.Create(&user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(int(user.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var deletedUser models.User
	err := db.Where("id = ?", user.ID).First(&deletedUser).Error
	assert.NotNil(t, err)
}

func TestUpdateUser(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	router := SetupRouter(db)

	// Create a user for testing
	user := models.User{Name: "John Doe", Email: "john.doe@example.com", Password: "password"}
	db.Create(&user)

	updatedUser := models.User{Name: "John Smith", Email: "john.doe@example.com", Password: "newpassword"}
	jsonValue, _ := json.Marshal(updatedUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/users/"+strconv.Itoa(int(user.ID)), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var userFromDB models.User
	db.Where("id = ?", user.ID).First(&userFromDB)
	assert.Equal(t, "John Smith", userFromDB.Name)

	err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte("newpassword"))
	assert.Nil(t, err)
}
