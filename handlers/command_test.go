package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

var commandsJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`

func TestGetCommands(test *testing.T) {
	// Setup
	router := echo.New()
	request := new(http.Request)
	recorder := httptest.NewRecorder()
	context := router.NewContext(standard.NewRequest(request, router.Logger()), standard.NewResponse(recorder, router.Logger()))
	context.SetPath("/command/:address")
	context.SetParamNames("address")
	context.SetParamValues("1.2.3.4")

	// Assertions
	if assert.NoError(test, GetCommands(context)) {
		assert.Equal(test, http.StatusOK, recorder.Code)
		assert.Equal(test, commandsJSON, recorder.Body.String())
	}
}
