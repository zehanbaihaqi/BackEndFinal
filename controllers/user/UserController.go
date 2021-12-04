package user

import (
	"net/http"
	"sim/configs"
	"sim/middlewares"
	"sim/models/response"
	"sim/models/users"

	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
	var user users.Users
	c.Bind(&user)

	result := configs.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = response.BaseResponse{
		http.StatusOK,
		"New user registered!",
		user,
	}
	return c.JSON(http.StatusOK, response)
}

func LoginController(c echo.Context) error {
	var user users.Users
	c.Bind(&user)

	if err := configs.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			http.StatusBadRequest,
			"Email dan Password tidak sesuai",
			nil,
		})
	}

	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			http.StatusBadRequest,
			"Password Kosong",
			nil,
		})
	}

	token, err := middlewares.GenerateTokenJWT(int(user.Id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			"Error ketika generate token JWT",
			nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		http.StatusOK,
		"succes",
		map[string]string{
			"token": token,
		},
	})
}
