package hdmiControl

import (
	"errors"
	"fmt"
	"time"

	"github.com/candiesdoodle/cec"
)

var hdmi *cec.Connection
var hdmiPort = 0

func init() {
	var err error

	hdmi, err = cec.Open("", "cec.go")
	if err != nil {
		fmt.Println(err)
	}
}

func SetPort(port int) {
	hdmiPort = port
}

func GetDeviceInfo(port int) cec.Device {
	devices := GetActiveDeviceList()

	for _, device := range devices {
		if device.LogicalAddress == port {
			return device
		}
	}

	return cec.Device{}
}

func GetActiveDeviceList() map[string]cec.Device {
	return hdmi.List()
}

func GetPowerStatus() string {
        result := hdmi.GetDevicePowerStatus(hdmiPort)
        fmt.Println("Getting Device Power for: ",hdmiPort)
        for i := 1; (i < 5 && result == ""); i++ {
          fmt.Println("Failed to Get PowerStatus -> I'll try it again", i)
          time.Sleep(1000 * time.Millisecond)
          result = hdmi.GetDevicePowerStatus(hdmiPort)
        }
        return result
}

func Power(state string) error {
	switch state {
	case "on":
		return hdmi.PowerOn(hdmiPort)
	case "off":
		return hdmi.Standby(hdmiPort)
	default:
		return errors.New("Invalid power state given.")
	}
}

func SetVolume(state string) error {
	switch state {
	case "up":
		return hdmi.VolumeUp()
	case "down":
		return hdmi.VolumeDown()
	case "mute":
		return hdmi.Mute()
	default:
		return errors.New("Invalid volume state given.")
	}
}

func Transmit(command string){
	hdmi.Transmit(command)
}

