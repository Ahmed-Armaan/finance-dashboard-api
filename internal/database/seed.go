package database

import (
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database/models"
)

func (db *DatabaseHolder) Seed(hashedPass string) error {
	superUser := &models.User{
		UserName: "super",
		Password: hashedPass,
		Role:     models.SuperAdmin,
	}

	if err := db.DB.Create(superUser).Error; err != nil {
		return err
	}
	return nil
}
