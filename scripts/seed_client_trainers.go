package scripts

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"ddx_hackathon_backend/models"

	"github.com/jinzhu/gorm"
)

func SeedClientTrainers(db *gorm.DB) {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

    var clients []models.User
    var trainers []models.User

    if err := db.Where("role = ?", "client").Find(&clients).Error; err != nil {
        log.Fatalf("Failed to fetch clients: %v", err)
    }
    
    if err := db.Where("role = ?", "trainer").Find(&trainers).Error; err != nil {
        log.Fatalf("Failed to fetch trainers: %v", err)
    }

    for _, client := range clients {
        numTrainers := rnd.Intn(3) // 0, 1, or 2 trainers
        assignedTrainers := make(map[uint]struct{})

        for i := 0; i < numTrainers; i++ {
            trainer := trainers[rnd.Intn(len(trainers))]
            if _, exists := assignedTrainers[trainer.ID]; !exists {
                clientTrainer := models.ClientTrainer{
                    ClientID:  client.ID,
                    TrainerID: trainer.ID,
                }

                if err := db.Create(&clientTrainer).Error; err != nil {
                    log.Fatalf("Failed to create client-trainer relationship: %v", err)
                }

                assignedTrainers[trainer.ID] = struct{}{}
            }
        }
    }

    fmt.Println("Client-trainer relationships added successfully")
}
