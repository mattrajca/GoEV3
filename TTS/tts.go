// Provides Text-to-Speech capabilities.
package TTS

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Speaks the given string of text at the specified volume and speed.
// `volume` ranges from 0 to 200.
// `speed` is measured in words per minute.
func SpeakV(text string, volume uint8, speed uint8) {
	if volume > 200 {
		volume = 200
	}

	c1 := exec.Command("espeak", "--stdout", "-s", strconv.FormatInt(int64(speed), 10), "-a", strconv.FormatInt(int64(volume), 10), fmt.Sprintf("\"%s\"", text))
	c2 := exec.Command("aplay")

	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
}

// Speaks the given string of text at maximum volume and with a moderate speed.
func Speak(text string) {
	SpeakV(text, 200, 130)
}
