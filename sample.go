package main

import (
	"fmt"
	"github.com/mattrajca/GoEV3/Button"
	"github.com/mattrajca/GoEV3/LED"
	"github.com/mattrajca/GoEV3/Motor"
	"github.com/mattrajca/GoEV3/Sensors"
	"github.com/mattrajca/GoEV3/Sound"
	"github.com/mattrajca/GoEV3/TTS"
	"time"
)

func playTune() {
	Sound.SetVolume(50)

	Sound.PlayToneAndRest(262, 180, 20)
	Sound.PlayToneAndRest(262, 180, 20)
	Sound.PlayToneAndRest(392, 180, 20)
	Sound.PlayToneAndRest(392, 180, 20)
	Sound.PlayToneAndRest(440, 180, 20)
	Sound.PlayToneAndRest(440, 180, 20)
	Sound.PlayToneAndRest(392, 380, 20)

	Sound.SetVolume(100)

	Sound.PlayToneAndRest(349, 180, 20)
	Sound.PlayToneAndRest(349, 180, 20)
	Sound.PlayToneAndRest(330, 180, 20)
	Sound.PlayToneAndRest(330, 180, 20)
	Sound.PlayToneAndRest(294, 180, 20)
	Sound.PlayToneAndRest(294, 180, 20)
	Sound.PlayToneAndRest(262, 400, 200)
}

func testTouchAndSound() {
	fmt.Println("Press the touch sensor (on port 2) to play a tune")

	sensor := Sensors.FindTouchSensor(Sensors.InPort2)
	sensor.Wait()

	playTune()
}

func testUltrasonic() {
	fmt.Println("Reading ultrasonic sensor (on port 2)")

	sensor := Sensors.FindUltrasonicSensor(Sensors.InPort2)
	value := sensor.ReadDistance()

	fmt.Printf("Ultrasonic distance reading: %d\n", value)
}

func testInfrared() {
	fmt.Println("Reading infrared sensor (on port 2)")

	sensor := Sensors.FindInfraredSensor(Sensors.InPort2)

	for n := 0; n < 10; n++ {
		value := sensor.ReadProximity()
		fmt.Printf("Infrared proximity reading: %d\n", value)

		time.Sleep(time.Second / 2)
	}
}

func testColor() {
	fmt.Println("Reading color values")

	sensor := Sensors.FindColorSensor(Sensors.InPort2)

	for n := 0; n < 10; n++ {
		color := sensor.ReadColor()
		rl := sensor.ReadReflectedLightIntensity()
		al := sensor.ReadAmbientLightIntensity()

		fmt.Printf("Color: %v\n", color)
		fmt.Printf("Reflected light: %v\n", rl)
		fmt.Printf("Ambient light: %v\n", al)

		time.Sleep(time.Second / 2)
	}
}

func testTTSAndButtons() {
	TTS.Speak("Hello! Press the following sequence of buttons:")

	fmt.Println("Press the following sequence of buttons:")
	fmt.Println("left, right, up, down, enter")

	Button.Wait(Button.Left)
	Button.Wait(Button.Right)
	Button.Wait(Button.Up)
	Button.Wait(Button.Down)
	Button.Wait(Button.Enter)
}

func testMotors() {
	fmt.Println("Make sure there is a motor plugged in to port A and press the center button to continue.")
	Button.Wait(Button.Enter)

	fmt.Println("Running motor A")
	Motor.Run(Motor.OutPortA, 40)

	for n := 0; n < 20; n++ {
		time.Sleep(time.Second / 2)

		fmt.Printf("Running motor A with speed %d and power %d\n", Motor.CurrentSpeed(Motor.OutPortA), Motor.CurrentPower(Motor.OutPortA))
	}

	Motor.Stop(Motor.OutPortA)
}

func testLEDs() {
	for n := 0; n < 12; n++ {
		if n%3 == 0 {
			LED.TurnOff(LED.Red, LED.Left)
			LED.TurnOff(LED.Red, LED.Right)

			LED.TurnOn(LED.Amber, LED.Left)
			LED.TurnOn(LED.Amber, LED.Right)
		} else if n%3 == 1 {
			LED.TurnOff(LED.Amber, LED.Left)
			LED.TurnOff(LED.Amber, LED.Right)

			LED.TurnOn(LED.Green, LED.Left)
			LED.TurnOn(LED.Green, LED.Right)
		} else {
			LED.TurnOff(LED.Green, LED.Left)
			LED.TurnOff(LED.Green, LED.Right)

			LED.TurnOn(LED.Red, LED.Left)
			LED.TurnOn(LED.Red, LED.Right)
		}

		time.Sleep(time.Second / 2)
	}
}

func dispatchCommand(selection int) {
	switch selection {
	case 0:
		{
			testTouchAndSound()
		}
	case 1:
		{
			testUltrasonic()
		}
	case 2:
		{
			testInfrared()
		}
	case 3:
		{
			testColor()
		}
	case 4:
		{
			testTTSAndButtons()
		}
	case 5:
		{
			testMotors()
		}
	case 6:
		{
			testLEDs()
		}
	}
}

func main() {
	selection := int(0)

	for {
		fmt.Println("\nWhat would you like to test? Press Escape to exit.")

		var entries = [7]string{
			"Touch sensor and sound",
			"Ultrasonic sensor",
			"Infrared sensor",
			"Color sensor",
			"Speech and brick buttons",
			"Motors",
			"LEDs",
		}

		for i, f := range entries {
			checked := " "

			if i == selection {
				checked = "X"
			}

			fmt.Printf("[%s] %s\n", checked, f)
		}

		fmt.Print("\n")

		button := Button.WaitAny()

		if button == Button.Down {
			if selection < len(entries)-1 {
				selection++
			}
		} else if button == Button.Up {
			if selection > 0 {
				selection--
			}
		} else if button == Button.Enter {
			dispatchCommand(selection)
		} else if button == Button.Escape {
			return
		}
	}
}
