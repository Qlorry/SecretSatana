package routes

import (
	"log"
	"net/http"
	"secret-satana/database"
	"secret-satana/models"

	"github.com/labstack/echo/v4"
)

func RegisterParticipateRoutes(e *echo.Echo) {
	e.POST("/participate", handleParticipate)
}

func handleParticipate(c echo.Context) error {
	username := c.Get("username")

	result := database.DB.
		Where("name = ?", username).
		Updates(models.User{Participates: true})

	if result.Error != nil {
		log.Println("Failed to update User:", username, result.Error.Error())
		return c.Render(http.StatusOK, "participate-result.html", map[string]interface{}{
			"Error": result.Error.Error(),
		})
	}
	return c.Render(http.StatusOK, "participate-result.html", nil)
}
