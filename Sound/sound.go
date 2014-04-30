// Provides sound playback capabilities.
package Sound

import (
	"github.com/mattrajca/GoEV3/utilities"
	"os/exec"
	"time"
)

// Plays the given wave file at system volume.
func Play(path string) {
	c1 := exec.Command("aplay", path)
	_ = c1.Run()
}

// Returns the current system volume in range [0, 100].
func CurrentVolume() uint8 {
	return utilities.ReadUInt8Value("/sys/devices/platform/snd-legoev3", "volume")
}

// Sets the system volume to the specified argument in range [0, 100].
func SetVolume(volume uint8) {
	if volume > 100 {
		volume = 100
	}

	utilities.WriteUIntValue("/sys/devices/platform/snd-legoev3", "volume", uint64(volume))
}

// Returns the frequency of the current playing tone.
func CurrentTone() uint32 {
	return utilities.ReadUInt32Value("/sys/devices/platform/snd-legoev3", "tone")
}

// Plays a tone at the given frequency for the given duration (in ms).
func PlayTone(freq uint32, duration uint64) {
	utilities.WriteUIntValue("/sys/devices/platform/snd-legoev3", "tone", uint64(freq))
	time.Sleep(time.Duration(duration) * time.Millisecond)
	utilities.WriteUIntValue("/sys/devices/platform/snd-legoev3", "tone", 0)
}

// Plays a tone at the given frequency for the given duration (in ms). Then sleeps for `rest` ms.
func PlayToneAndRest(freq uint32, duration uint64, rest uint64) {
	PlayTone(freq, duration)
	time.Sleep(time.Duration(rest) * time.Millisecond)
}
