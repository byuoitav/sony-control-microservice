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

	// commands, err := helpers.GetCommands(request.Address)
	// if err != nil {
	// 	return err
	// }

	// s := reflect.ValueOf(commands.Result)
	// s = reflect.ValueOf(s.Index(1))
	// for i := 0; i < s.Len(); i++ {
	// fmt.Println(s.Index(1))
	// commandArray := s.Index(1)
	// for i := range commandArray {
	// 	fmt.Println(commandArray[i])
	// }
	// }

	// for i := range commands.Result {
	// fmt.Println(i)
	// if strings.ToLower(request.Command) == strings.ToLower(j) {
	// 	fmt.Println("Found")
	// }
	// }

	// fmt.Printf("%+v\n", commands)

	response, err := helpers.SendCommand(request.Address, request.Command)
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
