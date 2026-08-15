package repository

import (
	"time"

	"github.com/Wenev/WeDone/backend/internal/auth/domain"
	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Email     string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Initial   string     `gorm:"type:varchar(10);uniqueIndex;not null"`
	Password  string     `gorm:"type:varchar(255)"`
	GoogleID  string     `gorm:"type:varchar(255);index"`
	Role      string     `gorm:"type:varchar(20);not null"`
	ManagerID *uuid.UUID `gorm:"type:uuid;index"`
	IsAdmin   bool       `gorm:"type:boolean;not null;default:false"`
	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
}

func ToDomain(m *UserModel) *domain.User {
	return &domain.User{
		ID:        m.ID,
		Email:     m.Email,
		Initial:   m.Initial,
		Password:  m.Password,
		GoogleID:  m.GoogleID,
		Role:      domain.Role(m.Role),
		ManagerID: m.ManagerID,
		IsAdmin:   m.IsAdmin,
	}
}

func FromDomain(d *domain.User) *UserModel {
	return &UserModel{
		ID:        d.ID,
		Email:     d.Email,
		Initial:   d.Initial,
		Password:  d.Password,
		GoogleID:  d.GoogleID,
		Role:      string(d.Role),
		ManagerID: d.ManagerID,
		IsAdmin:   d.IsAdmin,
	}
}
