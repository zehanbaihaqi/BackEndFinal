package routes

import (
	"sim/controllers/pemberitahuan"
	"sim/controllers/user"
	"sim/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutePemberitahuan() *echo.Echo {
	e := echo.New()

	ePemberitahuan := e.Group("/pemberitahuans")
	config := middleware.JWTConfig{
		Claims:     &middlewares.JwtCustomClaims{},
		SigningKey: []byte("kodeini"),
	}
	ePemberitahuan.Use(middleware.JWTWithConfig(config))
	//POST
	ePemberitahuan.POST("/", pemberitahuan.InsertNewPemberitahuanController)

	// GET
	ePemberitahuan.GET("/", pemberitahuan.GetPemberitahuanController)
	ePemberitahuan.GET("/:pemberitahuanId", pemberitahuan.GetDetailPemberitahuanController)

	//Auth
	e.POST("/register", user.RegisterController)
	e.POST("/login", user.LoginController)

	return e
}
