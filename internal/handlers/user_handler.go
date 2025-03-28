package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, []interface{}{
		map[string]string{"id": "1", "name": "Juan"},
		map[string]string{"id": "1", "name": "Pedro"},
	})

}
