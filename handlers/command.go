package handlers

import (
	"net/http"

	"github.com/byuoitav/sony-control-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func GetCommands(context echo.Context) error {
	response, err := helpers.GetCommands(context.Param("address"))
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func SendCommand(context echo.Context) error {
	response, err := helpers.SendCommand(context.Param("address"), context.Param("command"))
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
