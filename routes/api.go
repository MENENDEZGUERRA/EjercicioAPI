package routes

import (
	"EjercicioAPI/internal/handlers"
	"EjercicioAPI/internal/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	incidentRepo := repositories.NewIncidentRepository(db)
	incidentHandler := handlers.NewIncidentHandler(incidentRepo)

	api := e.Group("/api/v1")
	{
		api.POST("/incidents", incidentHandler.CreateIncident)
		api.GET("/incidents", incidentHandler.GetIncidents)
		api.GET("/incidents/:id", incidentHandler.GetIncident)
		api.PUT("/incidents/:id", incidentHandler.UpdateIncidentStatus)
		api.DELETE("/incidents/:id", incidentHandler.DeleteIncident)
	}
}
