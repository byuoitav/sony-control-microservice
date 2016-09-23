package handlers

import (
	"errors"
	"log"
	"strconv"

	"github.com/byuoitav/sony-control-microservice/helpers"
	"github.com/labstack/echo"
)

func PowerOn(context echo.Context) error {
	log.Printf("Powering on %s...", context.Param("address"))
	_, err := helpers.SendCommand(context.Param("address"), "wakeup")
	if err != nil {
		return err
	}
	log.Printf("Done.")
	return nil
}

func Standby(context echo.Context) error {
	log.Printf("Powering off %s...", context.Param("address"))
	_, err := helpers.SendCommand(context.Param("address"), "poweroff")
	if err != nil {
		return err
	}
	log.Printf("Done.")
	return nil
}

func SwitchInput(context echo.Context) error {
	log.Printf("Switching input for %s to %s ...", context.Param("address"), context.Param("port"))
	_, err := helpers.SendCommand(context.Param("address"), context.Param("port")) // Input must be hdmi1, hdmi2, etc.
	if err != nil {
		return err
	}
	log.Printf("Done.")
	return nil
}

func SetVolume(context echo.Context) error {
	address := context.Param("address")
	difference, err := strconv.Atoi(context.Param("difference"))
	if err != nil {
		return err
	}

	log.Printf("Setting volume for %s by %v...", address, difference)
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
	log.Printf("Done.")
	return nil
}

func CalibrateVolume(context echo.Context) error {
	address := context.Param("address")
	log.Printf("Calibrating volume for %s...", address)
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
	log.Printf("Done.")
	return nil
}

func VolumeUp(context echo.Context) error {
	log.Printf("Setting volume up for %s...", context.Param("address"))
	_, err := helpers.SendCommand(context.Param("address"), "volumeup")
	if err != nil {
		return err
	}

	return nil
}

func VolumeDown(context echo.Context) error {
	log.Printf("Setting volume down for %s...", context.Param("address"))
	_, err := helpers.SendCommand(context.Param("address"), "volumedown")
	if err != nil {
		return err
	}

	log.Printf("Done.")
	return nil
}
func VolumeUnmute(context echo.Context) error {
	log.Printf("Unmuting %s...", context.Param("address"))
	_, err := helpers.SendCommand(context.Param("address"), "mute")
	if err != nil {
		return err
	}

	return nil
}
func VolumeMute(context echo.Context) error {
	log.Printf("Muting %s...", context.Param("address"))
	_, err := helpers.SendCommand(context.Param("address"), "mute")
	if err != nil {
		return err
	}

	log.Printf("Done.")
	return nil
}

func BlankDisplay(context echo.Context) error {
	_, err := helpers.SendCommand(context.Param("address"), "pictureoff")
	if err != nil {
		return err
	}

	log.Printf("Done.")
	return nil
}

func UnblankDisplay(context echo.Context) error {
	log.Printf("Blanking %s...", context.Param("address"))
	_, err := helpers.SendCommand(context.Param("address"), "pictureoff")
	if err != nil {
		return err
	}
	log.Printf("Done.")
	return nil
}
