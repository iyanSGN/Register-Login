package app

import (
	"fmt"
	"goLANG/models"
	"goLANG/pkg/database"
)

func CreateUser(user models.MasterUser) (models.MasterUser, error) {
	db := database.GetDB()

	result := db.Create(&user)
	if result.Error != nil {
		fmt.Printf("Error create user: %v", result.Error)
	}

	return user, nil
}

func GetUser() ([]models.MasterUser, error) {
	db := database.GetDB()

	var users []models.MasterUser
	result := db.Find(&users)
	if result.Error != nil {
		fmt.Printf("error fetching user: %v", result.Error)
	}

	return users, nil
}


func UpdateUser(id int, userId models.MasterUser) error {
	db := database.GetDB()

	var user models.MasterUser
	result := db.First(&user, id)

	if result.Error != nil {
		return fmt.Errorf("error : %w", result.Error)
	}

	if userId.Name != "" {
		user.Name = userId.Name
	}

	if userId.Email != "" {
		user.Email = userId.Email
	}

	if userId.Password != "" {
		user.Password = userId.Password
	}

	updateResult := db.Save(&user)
	if updateResult.Error != nil {
		return fmt.Errorf("error saving updates: %w", updateResult.Error)
	}

	return nil
}

func DeleteUser (id int) error {
	db := database.GetDB()

	var user models.MasterUser
	result := db.Delete(&user, id)
	if result.Error != nil {
		return fmt.Errorf("error : %w", result.Error)
	}
	return nil
}
