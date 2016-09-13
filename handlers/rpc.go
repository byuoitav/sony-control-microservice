package handlers

import (
	"errors"
	"strconv"

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

func Standby(context echo.Context) error {
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

func SetVolume(context echo.Context) error {
	address := context.Param("address")
	difference, err := strconv.Atoi(context.Param("difference"))
	if err != nil {
		return err
	}
	//Setting volume up
	if difference > 0 {
		for i := 0; i < difference; i++ {
			_, err := helpers.SendCommand(address, "volumeup")
			if err != nil {
				return err
			}
		}
	} else if difference < 0 {
		for i := 0; i > difference; i-- {
			_, err := helpers.SendCommand(address, "volumedown")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CalibrateVolume(context echo.Context) error {
	address := context.Param("address")
	def, err := strconv.Atoi(context.Param("default"))
	if err != nil {
		return err
	}
	if def < 0 || def > 100 {
		return errors.New("Invalid default value, must be in range 0-100")
	}

	//drop volume to zero.
	for i := 0; i < 125; i++ {
		_, err := helpers.SendCommand(address, "volumedown")
		if err != nil {
			return err
		}
	}
	//set voulme to 25
	for i := 0; i < def; i++ {
		_, err := helpers.SendCommand(address, "volumeup")
		if err != nil {
			return err
		}
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
func VolumeUnmute(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), "mute")
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
