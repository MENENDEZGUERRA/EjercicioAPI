package handlers

import (
	"EjercicioAPI/internal/models"
	"EjercicioAPI/internal/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IncidentHandler struct {
	repo *repositories.IncidentRepository
}

func NewIncidentHandler(repo *repositories.IncidentRepository) *IncidentHandler {
	return &IncidentHandler{repo: repo}
}

func (h *IncidentHandler) CreateIncident(c echo.Context) error {
	var incident models.Incident
	if err := c.Bind(&incident); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := c.Validate(incident); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.repo.Create(&incident); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create incident"})
	}

	return c.JSON(http.StatusCreated, incident)
}

func (h *IncidentHandler) GetIncidents(c echo.Context) error {
	incidents, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve incidents"})
	}
	return c.JSON(http.StatusOK, incidents)
}

func (h *IncidentHandler) GetIncident(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	incident, err := h.repo.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve incident"})
	}

	if incident == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Incident not found"})
	}

	return c.JSON(http.StatusOK, incident)
}

func (h *IncidentHandler) UpdateIncidentStatus(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var update models.IncidentUpdate

	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := c.Validate(update); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.repo.UpdateStatus(uint(id), update.Status); err != nil {
		if err.Error() == "incident not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Incident not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update incident"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Status updated successfully"})
}

func (h *IncidentHandler) DeleteIncident(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.repo.Delete(uint(id)); err != nil {
		if err.Error() == "incident not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Incident not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete incident"})
	}

	return c.NoContent(http.StatusNoContent)
}
