package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/sony-control-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func GetCommands(context echo.Context) error {
	response, err := helpers.GetCommands(context.Param("address"))
	if err != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, err.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func SendCommand(context echo.Context) error {
	response, err := helpers.SendCommand(context.Param("address"), context.Param("command"))
	if err != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, err.Error())
		return nil
	}

	return context.JSON(http.StatusOK, response)
}

func FloodCommand(context echo.Context) error {
	count, err := strconv.Atoi(context.Param("count"))
	if err != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, err.Error())
		return nil
	}

	helpers.FloodCommand(context.Param("address"), context.Param("command"), count)

	return context.JSON(http.StatusOK, "Hello")
}
