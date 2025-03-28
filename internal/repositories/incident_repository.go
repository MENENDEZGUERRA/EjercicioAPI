package repositories

import (
	"EjercicioAPI/internal/models"
	"errors"

	"gorm.io/gorm"
)

type IncidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) *IncidentRepository {
	return &IncidentRepository{db: db}
}

func (r *IncidentRepository) Create(incident *models.Incident) error {
	return r.db.Create(incident).Error
}

func (r *IncidentRepository) GetAll() ([]models.Incident, error) {
	var incidents []models.Incident
	result := r.db.Find(&incidents)
	return incidents, result.Error
}

func (r *IncidentRepository) GetByID(id uint) (*models.Incident, error) {
	var incident models.Incident
	result := r.db.First(&incident, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &incident, result.Error
}

func (r *IncidentRepository) UpdateStatus(id uint, status models.IncidentStatus) error {
	result := r.db.Model(&models.Incident{}).Where("id = ?", id).Update("status", status)
	if result.RowsAffected == 0 {
		return errors.New("incident not found")
	}
	return result.Error
}

func (r *IncidentRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Incident{}, id)
	if result.RowsAffected == 0 {
		return errors.New("incident not found")
	}
	return result.Error
}
