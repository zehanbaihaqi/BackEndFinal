package pemberitahuan

import (
	"net/http"
	"sim/configs"
	"sim/models/pemberitahuan"
	"sim/models/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDetailPemberitahuanController(c echo.Context) error {
	pemberitahuanId, _ := strconv.Atoi(c.Param("pemberitahuanId"))
	var pemberitahuan []pemberitahuan.Pemberitahuans

	dataPemberitahuans := configs.DB.First(&pemberitahuan, pemberitahuanId)

	if dataPemberitahuans.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			dataPemberitahuans.Error.Error(),
			nil,
		})
	}

	var response = response.BaseResponse{
		http.StatusOK,
		"success",
		pemberitahuan,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateDataPemberitahuan(c echo.Context) error {
	pemberitahuanId, _ := strconv.Atoi(c.Param("pemberitahuanId"))

	var data = pemberitahuan.Pemberitahuans{}

	var response = response.BaseResponse{
		pemberitahuanId,
		"success",
		data,
	}

	return c.JSON(http.StatusOK, response)
}

func GetPemberitahuanController(c echo.Context) error {
	var pemberitahuan []pemberitahuan.Pemberitahuans

	dataPemberitahuans := configs.DB.Find(&pemberitahuan)

	if dataPemberitahuans.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			dataPemberitahuans.Error.Error(),
			nil,
		})
	}

	var response = response.BaseResponse{
		http.StatusOK,
		"success",
		pemberitahuan,
	}

	return c.JSON(http.StatusOK, response)
}

func InsertNewPemberitahuanController(c echo.Context) error {
	var pemberitahuan pemberitahuan.Pemberitahuans
	c.Bind(&pemberitahuan)

	newPemberitahuan := configs.DB.Create(&pemberitahuan)

	if newPemberitahuan.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			newPemberitahuan.Error.Error(),
			nil,
		})
	}

	var response = response.BaseResponse{
		http.StatusOK,
		"success",
		pemberitahuan,
	}

	return c.JSON(http.StatusOK, response)
}
