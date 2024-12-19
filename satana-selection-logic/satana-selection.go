package satana_selection

import (
	"fmt"
	"slices"

	"math/rand/v2"
	"secret-satana/database"
	"secret-satana/models"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func ReselectSatanas() error {
	var oldPairs []models.SatanaSelection
	result := database.DB.Find(&oldPairs)
	if result.Error != nil {
		return result.Error
	}
	database.DB.Delete(&oldPairs)

	var allUsers []models.User
	if result := database.DB.Where("participates = ?", true).Find(&allUsers); result.Error != nil {
		return result.Error
	}

	if len(allUsers) < 2 {
		return fmt.Errorf("Not enough users to select satanas")
	}

	if len(allUsers)%2 != 0 {
		return fmt.Errorf("Number of users must be even")
	}

	var selectedSatanas = make([]models.SatanaSelection, 0, len(allUsers)/2)
	usersWithoutSatana := append(make([]models.User, 0, len(allUsers)), allUsers...)

	for i := 0; i < len(allUsers); i++ {
		currentUsersWithoutSatana := append(make([]models.User, 0, len(usersWithoutSatana)), usersWithoutSatana...)

		currentUsersWithoutSatana = slices.DeleteFunc(currentUsersWithoutSatana, func(user models.User) bool {
			return user.ID == allUsers[i].ID
		})

		secterSatanaIndex := randRange(0, len(currentUsersWithoutSatana))
		secterSatanaId := currentUsersWithoutSatana[secterSatanaIndex].ID

		selectedSatanas = append(selectedSatanas, models.SatanaSelection{UserOneID: allUsers[i].ID, UserTwoID: secterSatanaId})
		usersWithoutSatana = slices.DeleteFunc(usersWithoutSatana, func(user models.User) bool {
			return user.ID == secterSatanaId
		})
	}

	if err := database.DB.Create(&selectedSatanas).Error; err != nil {
		return err
	}

	return nil
}
