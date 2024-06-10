package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"ddx_hackathon_backend/handlers"
	"ddx_hackathon_backend/models"
)

// Mock database
func setupTestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	return db
}

// Mock router
func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/users", func(c *gin.Context) {
		handlers.GetUsers(c, db)
	})
	router.POST("/users", func(c *gin.Context) {
		handlers.CreateUser(c, db)
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		handlers.DeleteUser(c, db)
	})
	router.PATCH("/users/:id", func(c *gin.Context) {
		handlers.UpdateUser(c, db)
	})
	return router
}

func TestGetUsers(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	router := setupRouter(db)

	// Create a user for testing
	db.Create(&models.User{Name: "John Doe", Email: "john.doe@example.com", Password: "password"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	router := setupRouter(db)

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
	db := setupTestDB()
	defer db.Close()

	router := setupRouter(db)

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
	db := setupTestDB()
	defer db.Close()

	router := setupRouter(db)

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
