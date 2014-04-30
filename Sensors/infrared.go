package Sensors

import (
	"fmt"
	"github.com/mattrajca/GoEV3/utilities"
)

// Infrared sensor type.
type InfraredSensor struct {
	port InPort
}

// Provides access to an infrared sensor at the given port.
func FindInfraredSensor(port InPort) *InfraredSensor {
	findSensor(port, TypeInfrared)

	s := new(InfraredSensor)
	s.port = port

	return s
}

// Reads the proximity value (in range 0 - 100) reported by the infrared sensor at the given port. A value of 100 corresponds to a range of approximately 70 cm.
func (self *InfraredSensor) ReadProximity() uint8 {
	snr := findSensor(self.port, TypeInfrared)

	path := fmt.Sprintf("/sys/class/msensor/%s", snr)
	utilities.WriteStringValue(path, "mode", "IR-PROX")
	value := utilities.ReadUInt8Value(path, "value0")

	return value
}
