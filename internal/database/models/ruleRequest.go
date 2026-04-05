package models

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserId        uuid.UUID `gorm:"type:uuid;not null"`
	user          User      `gorm:"foreignKey:UserId"` // foreign key for the user who created the request
	RequestedRole UserRole  `gorm:"not null"`
	Resolved      bool      `gorm:"not null;default:false"`
	CreatedAt     time.Time
}
