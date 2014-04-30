package Sensors

import (
	"fmt"
	"github.com/mattrajca/GoEV3/utilities"
)

// Color sensor type.
type ColorSensor struct {
	port InPort
}

// Provides access to a color sensor at the given port.
func FindColorSensor(port InPort) *ColorSensor {
	findSensor(port, TypeColor)

	s := new(ColorSensor)
	s.port = port

	return s
}

// Constants for color values.
type Color uint8

const (
	None   Color = 0
	Black        = 1
	Blue         = 2
	Green        = 3
	Yellow       = 4
	Red          = 5
	White        = 6
	Brown        = 7
)

func (self Color) String() string {
	switch self {
	case Black:
		return "Black"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Red:
		return "Red"
	case White:
		return "White"
	case Brown:
		return "Brown"
	default:
		return "None"
	}
}

// Reads one of seven color values.
func (self *ColorSensor) ReadColor() Color {
	snr := findSensor(self.port, TypeColor)

	path := fmt.Sprintf("/sys/class/msensor/%s", snr)
	utilities.WriteStringValue(path, "mode", "COL-COLOR")
	value := utilities.ReadUInt8Value(path, "value0")

	return Color(value)
}

// Reads the reflected light intensity in range [0, 100].
func (self *ColorSensor) ReadReflectedLightIntensity() uint8 {
	snr := findSensor(self.port, TypeColor)

	path := fmt.Sprintf("/sys/class/msensor/%s", snr)
	utilities.WriteStringValue(path, "mode", "COL-REFLECT")
	value := utilities.ReadUInt8Value(path, "value0")

	return value
}

// Reads the ambient light intensity in range [0, 100].
func (self *ColorSensor) ReadAmbientLightIntensity() uint8 {
	snr := findSensor(self.port, TypeColor)

	path := fmt.Sprintf("/sys/class/msensor/%s", snr)
	utilities.WriteStringValue(path, "mode", "COL-AMBIENT")
	value := utilities.ReadUInt8Value(path, "value0")

	return value
}
