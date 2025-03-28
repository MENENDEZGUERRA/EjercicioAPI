package models

import (
	"time"

	"gorm.io/gorm"
)

type IncidentStatus string

const (
	StatusPending   IncidentStatus = "pendiente"
	StatusInProcess IncidentStatus = "en proceso"
	StatusResolved  IncidentStatus = "resuelto"
)

type Incident struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey"`
	Reporter    string         `json:"reporter" gorm:"not null" validate:"required"`
	Description string         `json:"description" gorm:"not null" validate:"required,min=10"`
	Status      IncidentStatus `json:"status" gorm:"default:'pendiente';not null"`
	CreatedAt   time.Time      `json:"created_at"`
}

type IncidentUpdate struct {
	Status IncidentStatus `json:"status" validate:"required,oneof=pendiente 'en proceso' resuelto"`
}