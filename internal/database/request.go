package database

import "github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database/models"

func (db *DatabaseHolder) ListRequests(offset int) ([]models.Request, error) {
	var requests []models.Request

	if err := db.DB.Select("*", requests).
		Limit(10).
		Offset(offset).Error; err != nil {
		return requests, err
	}

	return requests, nil
}
