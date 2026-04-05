package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	Viewer     UserRole = "viewer"
	Analyst    UserRole = "analyst"
	Admin      UserRole = "admin"
	SuperAdmin UserRole = "super-admin"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserName  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      UserRole  `gorm:"not null"`
	IsDeleted bool      `gorm:"not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
