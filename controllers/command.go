package controllers

import (
	"fmt"
	"net/http"

	"github.com/byuoitav/sony-control-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

type request struct {
	Address string `json:"address"`
	Command string `json:"command"`
}

func Command(context echo.Context) error {
	request := request{}
	err := context.Bind(&request)
	if err != nil {
		return jsonresp.Create(context, http.StatusBadRequest, err.Error())
	}

	fmt.Printf("%+v\n", request)

	response, err := helpers.SendCommand(request.Address, request.Command)
	if err != nil {
		return jsonresp.Create(context, http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
