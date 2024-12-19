package routes

import (
	"fmt"
	"secret-satana/database"
	"secret-satana/models"
	"strings"

	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("my-secret-key")

func RegisterLoginRoutes(e *echo.Echo) {
	e.GET("/login", showLoginPage)
	e.GET("/register", showRegisterPage)
	e.POST("/login", handleLogin)
	e.POST("/register", handleRegister)
}

func showRegisterPage(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}

func showLoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func handleLogin(c echo.Context) error {
	payload := new(models.LoginPayload)
	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	var user models.User
	result := database.DB.Where("name = ?", payload.Username).First(&user)

	if result.Error != nil {
		return c.String(http.StatusBadRequest, "Could not find user")
	}

	if user.Password != payload.Password {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}

	tokenString, err := createToken(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	setTokenInCookie(tokenString, c)
	return c.Redirect(http.StatusSeeOther, "/index")
}

func handleRegister(c echo.Context) error {
	payload := new(models.RegisterPayload)
	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	if payload.Password != payload.RepeatPassword {
		return c.String(http.StatusBadRequest, "Passwords do not match")
	}

	var user models.User
	result := database.DB.Where("name = ?", payload.Username).First(&user)

	if result.Error == nil {
		return c.String(http.StatusBadRequest, "User already exists")
	}

	user = models.User{Name: payload.Username, Password: payload.Password}
	result = database.DB.Create(&user)

	if result.Error != nil {
		return c.String(http.StatusBadRequest, "Error creating user:"+result.Error.Error())
	} else {
		log.Println("User created successfully:", user)
	}

	tokenString, err := createToken(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	setTokenInCookie(tokenString, c)

	return c.Redirect(http.StatusSeeOther, "/index")
}

func createToken(user models.User) (string, error) {
	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("could not create token: %w", err)
	}

	return tokenString, nil
}

func setTokenInCookie(token string, c echo.Context) error {
	// Return token in cookie
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 1)
	c.SetCookie(cookie)
	return nil
}

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Request().URL.Path, "/public") ||
			c.Request().URL.Path == "/login" ||
			c.Request().URL.Path == "/register" {
			return next(c)
		}

		cookie, err := c.Cookie("token")
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		// Parse the token
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		// Set username in context
		claims := token.Claims.(jwt.MapClaims)
		c.Set("username", claims["username"])

		return next(c)
	}
}
