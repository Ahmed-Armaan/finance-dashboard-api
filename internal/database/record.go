package database

import (
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database/models"
	"github.com/google/uuid"
)

func (db *DatabaseHolder) CreateRoleRequest(userID uuid.UUID, requestedRole models.UserRole) error {
	request := &models.Request{
		UserId:        userID,
		RequestedRole: requestedRole,
	}
	return db.DB.Create(request).Error
}
