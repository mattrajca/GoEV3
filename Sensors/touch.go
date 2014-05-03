package Sensors

import (
	"fmt"
	"github.com/mattrajca/GoEV3/utilities"
	"time"
)

// Touch sensor type.
type TouchSensor struct {
	port InPort
}

// Provides access to a touch sensor at the given port.
func FindTouchSensor(port InPort) *TouchSensor {
	findSensor(port, TypeTouch)

	s := new(TouchSensor)
	s.port = port

	return s
}

// Waits for the touch sensor to be pressed.
func (self *TouchSensor) Wait() {
	snr := findSensor(self.port, TypeTouch)
	path := fmt.Sprintf("/sys/class/msensor/%s", snr)

	for {
		value := utilities.ReadUInt8Value(path, "value0")

		if value == 1 {
			return
		}

		time.Sleep(time.Millisecond * 50)
	}
}
