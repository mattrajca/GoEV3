// Provides APIs for interacting with EV3's LEDs.
package LED

import (
	"fmt"
	"github.com/mattrajca/GoEV3/utilities"
	"log"
	"os"
)

// Constants for the LED positions (left and right).
type Position string

const (
	Left  Position = "left"
	Right          = "right"
)

// Constants for the LED colors.
type Color string

const (
	Green Color = "green"
	Red         = "red"
	Amber       = "amber"
)

func findFilename(color Color, position Position) string {
	if color == Amber {
		log.Fatal("Amber colors must be decomposed into green and red")
	}

	filename := fmt.Sprintf("/sys/class/leds/ev3:%s:%s", string(color), string(position))

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal("Cannot find the LED interface")
	}

	return filename
}

// Turns on the given LED with the specified color.
func TurnOn(color Color, position Position) {
	if color == Amber {
		utilities.WriteIntValue(findFilename(Green, position), "brightness", 1)
		utilities.WriteIntValue(findFilename(Red, position), "brightness", 1)
	} else {
		utilities.WriteIntValue(findFilename(color, position), "brightness", 1)
	}
}

// Turns off the given LED with the specified color.
func TurnOff(color Color, position Position) {
	if color == Amber {
		utilities.WriteIntValue(findFilename(Green, position), "brightness", 0)
		utilities.WriteIntValue(findFilename(Red, position), "brightness", 0)
	} else {
		utilities.WriteIntValue(findFilename(color, position), "brightness", 0)
	}
}
