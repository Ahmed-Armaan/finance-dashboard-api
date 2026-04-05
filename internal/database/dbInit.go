package database

import (
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database/models"
	"github.com/google/uuid"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseStore interface {
	Seed(password string) error
	GetUser(username string) (*models.User, error)
	AddUser(username, hashedPassword string, role models.UserRole) (*models.User, error)
	UpdateUserRole(userID uuid.UUID, role models.UserRole) error
	CreateRoleRequest(userID uuid.UUID, requestedRole models.UserRole) error
	ListRequests(offset int) ([]models.Request, error)
}

type DatabaseHolder struct {
	DB *gorm.DB
}

func DbInit() (DatabaseStore, error) {
	dsn := os.Getenv("DATABASE_URL")
	databaseStore := &DatabaseHolder{}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return databaseStore, err
	}

	if err := db.AutoMigrate(&models.User{},
		&models.Record{},
		&models.Request{}); err != nil {
		return databaseStore, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return databaseStore, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	databaseStore.DB = db
	return databaseStore, nil
}
