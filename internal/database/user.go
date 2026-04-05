package database

import (
	"errors"

	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrUserInvalid  = errors.New("user does not exist")
	ErrUserExists   = errors.New("username already taken")
	ErrUserInactive = errors.New("user is inactive")
)

func (db *DatabaseHolder) GetUser(username string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("user_name = ? AND is_deleted = false", username).
		First(&user).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, ErrUserInvalid
		}
		return &user, err
	}
	return &user, nil
}

func (db *DatabaseHolder) AddUser(username, hashedPassword string, role models.UserRole) (*models.User, error) {
	var existing models.User
	err := db.DB.Where("user_name = ?", username).First(&existing).Error
	if err == nil {
		return nil, ErrUserExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user := &models.User{
		UserName: username,
		Password: hashedPassword,
		Role:     role,
	}

	if err := db.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DatabaseHolder) UpdateUserRole(userID uuid.UUID, role models.UserRole) error {
	result := db.DB.Model(&models.User{}).
		Where("id = ? AND is_deleted = false", userID).
		Update("role", role)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrUserInvalid
	}
	return nil
}
