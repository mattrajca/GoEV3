// Provides APIs for interacting with EV3's sensors.
package Sensors

import (
	"fmt"
	"github.com/mattrajca/GoEV3/utilities"
	"io/ioutil"
	"log"
	"strings"
)

// Constants for input ports.
type InPort string

const (
	InPort1 InPort = "in1"
	InPort2        = "in2"
	InPort3        = "in3"
	InPort4        = "in4"
)

// Constants for sensor types.
type Type string

const (
	TypeTouch      Type = "ev3-uart-16"
	TypeColor           = "ev3-uart-29"
	TypeUltrasonic      = "ev3-uart-30"
	TypeInfrared        = "ev3-uart-33"
)

func (self Type) String() string {
	switch self {
	case TypeTouch:
		return "touch"
	case TypeColor:
		return "color"
	case TypeUltrasonic:
		return "ultrasonic"
	case TypeInfrared:
		return "infrared"
	default:
		return "unknown"
	}
}

func findSensor(port InPort, t Type) string {
	sensors, _ := ioutil.ReadDir("/sys/class/msensor")

	for _, item := range sensors {
		if strings.HasPrefix(item.Name(), "sensor") {
			sensorPath := fmt.Sprintf("/sys/class/msensor/%s", item.Name())
			portr := utilities.ReadStringValue(sensorPath, "port_name")

			if InPort(portr) == port {
				typer := utilities.ReadStringValue(sensorPath, "name")
        
				if Type(typer) == t {
					return item.Name()
				}
			}
		}
	}

	log.Fatalf("Could not find %v sensor on port %v\n", t, port)

	return ""
}
