package helpers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

var postBody = []byte(`{}`)
var serverResponse = `{"response": "Poots"}`

func TestPostHTTP(test *testing.T) {
	// Setup
	router := echo.New()
	request := new(http.Request)
	recorder := httptest.NewRecorder()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()

	context := router.NewContext(standard.NewRequest(request, router.Logger()), standard.NewResponse(recorder, router.Logger()))
	context.SetPath("/command")
	context.SetParamNames("address")
	context.SetParamValues(server.URL[7:]) // Trim the `http://` off the server's URL

	// Assertions
	if assert.NoError(test, PostHTTP(server.URL[7:], postBody)) {
		assert.Equal(test, http.StatusOK, recorder.Code)
		assert.Equal(test, response, recorder.Body.String())
	}
}
