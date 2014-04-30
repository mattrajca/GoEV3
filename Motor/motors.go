// Provides APIs for interacting with EV3's motors.
package Motor

import (
	"fmt"
	"github.com/mattrajca/GoEV3/utilities"
	"log"
	"os"
)

// Constants for output ports.
type OutPort string

const (
	OutPortA OutPort = "A"
	OutPortB         = "B"
	OutPortC         = "C"
	OutPortD         = "D"
)

func findFilename(port OutPort) string {
	filename := fmt.Sprintf("/sys/class/tacho-motor/out%s:motor:tacho", string(port))

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal("Cannot connect to motor %s", port)
	}

	return filename
}

// Runs the motor at the given port.
// speed ranges from -100 to 100 and indicates the target speed of the motor, with negative values indicating reverse motion. Depending on the environment, the actual speed of the motor may be lower than the target speed.
func Run(port OutPort, speed int16) {
	if speed > 100 || speed < -100 {
		log.Fatal("The speed must be in range [-100, 100]")
	}

	utilities.WriteIntValue(findFilename(port), "run", 1)
	utilities.WriteIntValue(findFilename(port), "speed_setpoint", int64(speed))
}

// Stops the motor at the given port.
func Stop(port OutPort) {
	utilities.WriteIntValue(findFilename(port), "run", 0)
}

// Reads the actual speed of the motor at the given port.
func CurrentSpeed(port OutPort) int16 {
	return utilities.ReadInt16Value(findFilename(port), "speed")
}

// Reads the actual power of the motor at the given port.
func CurrentPower(port OutPort) int16 {
	return utilities.ReadInt16Value(findFilename(port), "power")
}

// Enables regulation mode, causing the motor at the given port to compensate for any resistance and maintain its target speed.
func EnableRegulationMode(port OutPort) {
	utilities.WriteStringValue(findFilename(port), "regulation_mode", "on")
}

// Disables regulation mode. Regulation mode is off by default.
func DisableRegulationMode(port OutPort) {
	utilities.WriteStringValue(findFilename(port), "regulation_mode", "off")
}

// Enables brake mode, causing the motor at the given port to brake to stops.
func EnableBrakeMode(port OutPort) {
	utilities.WriteStringValue(findFilename(port), "brake_mode", "on")
}

// Disables brake mode, causing the motor at the given port to coast to stops. Brake mode is off by default.
func DisableBrakeMode(port OutPort) {
	utilities.WriteStringValue(findFilename(port), "brake_mode", "off")
}
