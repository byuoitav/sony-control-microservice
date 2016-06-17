package handlers

import (
	"net/http"

	"github.com/byuoitav/sony-control-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

type request struct {
	Address string `json:"address"`
	Command string `json:"command"`
}

func GetCommands(context echo.Context) error {
	response, err := helpers.GetCommands(context.Param("address"))
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func SendCommand(context echo.Context) error {
	request := request{}
	err := context.Bind(&request)
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, err.Error())
	}

	response, err := helpers.SendCommand(request.Address, request.Command)
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
