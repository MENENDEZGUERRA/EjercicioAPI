package models

import (
	"time"

	"gorm.io/gorm"
)

// User representa la entidad de usuario en la base de datos y en el sistema
type User struct {
	gorm.Model
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	// Campos personalizados
	FirstName string     `json:"first_name" gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	LastName  string     `json:"last_name" gorm:"type:varchar(100);not null" validate:"required,min=2,max=100"`
	Email     string     `json:"email" gorm:"type:varchar(100);uniqueIndex;not null" validate:"required,email"`
	BirthDate *time.Time `json:"birth_date,omitempty" validate:"omitempty,lte"`
	Password  string     `json:"-" gorm:"type:varchar(255);not null" validate:"required,min=8,max=64"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
}

// UserResponse estructura para respuestas API (excluye campos sensibles)
type UserResponse struct {
	ID        uint       `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
}

// ToResponse convierte User a UserResponse
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		BirthDate: u.BirthDate,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
	}
}
