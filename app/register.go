package app

import (
	"fmt"
	"goLANG/models"
	"goLANG/pkg/database"
)

func CreateUser(user models.MasterUser) (models.MasterUser, error) {
	db := database.GetDB()

	err := db.Create(&user).Error
	if err != nil {
		fmt.Printf("Error create user: %v", err.Error())
	}

	return user, nil
}


func GetUser() ([]models.MasterUser, error) {
	db := database.GetDB()

	var users []models.MasterUser
	result := db.
	Preload("MasterDepartment").
	Find(&users)
	if result.Error != nil {
		fmt.Println("error fetching user: %w", result.Error)
	}

	return users, nil
}

func GetId(id int) ([]models.MasterUser, error)  {
	db := database.GetDB()

	var user []models.MasterUser
	result := db.
		Preload("MasterDepartment").
		First(&user, id)
	if result.Error != nil {
		return nil, fmt.Errorf("error retrieving user: %w", result.Error) 
	}

	getUser := make([]models.MasterUser, len(user))
	for i, user := range user {
		getUser[i] = models.MasterUser(user)
	}

	return getUser, nil

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

func DeleteUser(id int) error {
	db := database.GetDB()

	var user models.MasterUser
	result := db.Delete(&user, id)
	if result.Error != nil {
		return fmt.Errorf("error : %w", result.Error)
	}
	return nil
}

func GetUserByDepId(department_id int) ([]models.MasterUser, error) {
	db := database.GetDB()

	var department []models.MasterUser
	result := db.
		Preload("MasterDepartment").
		Where("department_id = ?",department_id).
		Find(&department)
	if result.Error != nil {
		return nil, fmt.Errorf("error retrieving role: %w", result.Error)
	}

	getDepartment := make([]models.MasterUser, len(department))
	for i, department := range department {
		getDepartment[i] = models.MasterUser(department)
	}

	return getDepartment, nil

}

