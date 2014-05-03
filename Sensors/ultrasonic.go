package Sensors

import (
	"fmt"
	"github.com/mattrajca/GoEV3/utilities"
)

// Ultrasonic sensor type.
type UltrasonicSensor struct {
	port InPort
}

// Provides access to an ultrasonic sensor at the given port.
func FindUltrasonicSensor(port InPort) *UltrasonicSensor {
	findSensor(port, TypeUltrasonic)

	s := new(UltrasonicSensor)
	s.port = port

	return s
}

// Reads the distance (in centimeters) reported by the ultrasonic sensor.
func (self *UltrasonicSensor) ReadDistance() uint16 {
	snr := findSensor(self.port, TypeUltrasonic)

	path := fmt.Sprintf("/sys/class/msensor/%s", snr)
	utilities.WriteStringValue(path, "mode", "US-SI-CM")
	value := utilities.ReadUInt16Value(path, "value0")

	return (value / 10)
}

// Looks for other nearby ultrasonic sensors and returns true if one is found.
func (self *UltrasonicSensor) Listen() bool {
	snr := findSensor(self.port, TypeUltrasonic)

	path := fmt.Sprintf("/sys/class/msensor/%s", snr)
	utilities.WriteStringValue(path, "mode", "US-LISTEN")
	value := utilities.ReadUInt8Value(path, "value0")

	if value == 1 {
		return true
	}

	return false
}
