package routes

import (
	"log"
	"net/http"

	configuration "secret-satana/configs"
	"secret-satana/database"
	"secret-satana/models"

	"github.com/labstack/echo/v4"
)

func RegisterIndexRoutes(e *echo.Echo) {
	e.GET("/", handleIndexPage)
	e.GET("/index", handleIndexPage)
}

func handleIndexPage(c echo.Context) error {
	username := c.Get("username")

	var user models.User
	if result := database.DB.Where("name = ?", username).First(&user); result.Error != nil {
		log.Println("Failed to get User:", username, result.Error.Error())
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	var pair models.SatanaSelection
	if configuration.SatanaSelected {
		if result := database.DB.Preload("UserTwo").Where("user_one_id = ?", user.ID).First(&pair); result.Error != nil {
			log.Println("Failed to get User:", username, result.Error.Error())
		}
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"User":           user,
		"SatanaSelected": configuration.SatanaSelected,
		"SelectedSatana": pair.UserTwo,
	})
}
