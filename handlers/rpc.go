package handlers

import (
	"github.com/byuoitav/sony-control-microservice/helpers"
	"github.com/labstack/echo"
)

func PowerOn(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), "tvpower")
	if err != nil {
		return err
	}

	return nil
}

func PowerOff(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), "poweroff")
	if err != nil {
		return err
	}

	return nil
}

func SwitchInput(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), context.Param("port")) // Input must be hdmi1, hdmi2, etc.
	if err != nil {
		return err
	}

	return nil
}

func VolumeUp(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), "volumeup")
	if err != nil {
		return err
	}

	return nil
}

func VolumeDown(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), "volumedown")
	if err != nil {
		return err
	}

	return nil
}

func VolumeMute(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), "mute")
	if err != nil {
		return err
	}

	return nil
}

func BlankDisplay(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), "pictureoff")
	if err != nil {
		return err
	}

	return nil
}

func UnblankDisplay(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), "pictureoff")
	if err != nil {
		return err
	}

	return nil
}
