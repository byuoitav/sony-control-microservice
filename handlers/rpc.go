package handlers

import "github.com/labstack/echo"

func PowerOn(context echo.Context) error {
	return nil
}

func PowerOff(context echo.Context) error {
	return nil
}

func SwitchInput(context echo.Context) error {
	context.Param("port")

	return nil
}

func ChangeVolume(context echo.Context) error {
	context.Param("level")

	return nil
}

func BlankDisplay(context echo.Context) error {
	return nil
}

func UnblankDisplay(context echo.Context) error {
	return nil
}
