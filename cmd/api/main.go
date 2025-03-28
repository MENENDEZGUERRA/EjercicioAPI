package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Message define la estructura que usaremos para el endpoint POST.
type Message struct {
	Text string `json:"text"`
}

func main() {
	// Esto sirve para crear una nueva instancia de Echo.
	e := echo.New()

	// Endpoint raíz: Devuelve un mensaje de bienvenida.
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Bienvenido a mi API",
		})
	})

	// Endpoint GET
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "¡Hola, mundo!",
		})
	})

	// Endpoint POST
	e.POST("/echo", func(c echo.Context) error {
		m := new(Message)
		if err := c.Bind(m); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Datos inválidos",
			})
		}
		return c.JSON(http.StatusOK, m)
	})

	// Inicia el servidor en localhost:1323 <----- En Puesto 1323!!!!!!!
	e.Logger.Fatal(e.Start(":1323"))
}
