package controllers

import (
	"net/http"

	"github.com/byuoitav/sony-control-microservice/packages/ircc"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

func Capability(context echo.Context) error {
	response, err := ircc.GetCapability(context.Param("address"))
	if err != nil {
		return jsonresp.Create(context, http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
