package app

import (
	"fmt"
	"goLANG/models"
	"goLANG/pkg/database"
)

func GetUserByEmail(email string) (*models.MasterUser, error) {
	db := database.GetDB()
	akun := &models.MasterUser{}

	result := db.Where("email = $1", email).First(akun)
	if result.Error != nil {
		return nil, fmt.Errorf("error fetching akun: %w", result.Error)
	}
	return akun, nil
}
