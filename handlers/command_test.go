package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

var commandsJSON = `{"id":10,"result":[{"bundled":true,"type":"IR_REMOTE_BUNDLE_TYPE_UC"},[{"name":"Num1","value":"AAAAAQAAAAEAAAAAAw=="},{"name":"Num2","value":"AAAAAQAAAAEAAAABAw=="}]]}`

func TestGetCommands(test *testing.T) {
	// Setup
	router := echo.New()
	request := new(http.Request)
	recorder := httptest.NewRecorder()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(writer, commandsJSON)
	}))
	defer server.Close()

	context := router.NewContext(standard.NewRequest(request, router.Logger()), standard.NewResponse(recorder, router.Logger()))
	context.SetPath("/command/:address")
	context.SetParamNames("address")
	context.SetParamValues(server.URL[7:]) // Trim the `http://` off the server's URL

	// Assertions
	if assert.NoError(test, GetCommands(context)) {
		assert.Equal(test, http.StatusOK, recorder.Code)
		assert.Equal(test, commandsJSON, recorder.Body.String())
	}
}
