// Provides APIs for interacting with EV3's buttons.
package Button

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"sync"
	"time"
)

var bSync = &sync.Mutex{}

// Constants for the brick's buttons.
type Kind uint8

const (
	Up     Kind = 103
	Down        = 108
	Left        = 105
	Right       = 106
	Enter       = 28
	Escape      = 1
)

func findFilename() string {
	filename := "/dev/input/by-path/platform-gpio-keys.0-event"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal("Cannot find keys file")
	}

	return filename
}

// Waits for the given button to be pressed.
func Wait(kind Kind) {
	for {
		bSync.Lock()

		f, _ := os.OpenFile(findFilename(), os.O_RDONLY, 0644)
		b := make([]byte, 16)
		f.Read(b)
		f.Close()

		codeData := b[10:11]
		valueData := b[12:13]

		buf := bytes.NewBuffer(codeData)
		var code int8
		binary.Read(buf, binary.BigEndian, &code)

		buf2 := bytes.NewBuffer(valueData)
		var value int8
		binary.Read(buf2, binary.BigEndian, &value)

		bSync.Unlock()

		if value == 0 && code == int8(kind) {
			break
		}
	}
}

// Asynchronously waits for the given button to be pressed. `kind` is passed back to the given channel.
func WaitOnChannel(kind Kind, queue chan Kind) {
	go func() {
		Wait(kind)
		queue <- kind
	}()
}

// Waits for the given button to be pressed in a time window.
func WaitWithTimeout(kind Kind, t time.Duration) {
	c1 := make(chan Kind, 1)

	WaitOnChannel(kind, c1)

	select {
	case <-c1:
		{

		}
	case <-time.After(t):
		{

		}
	}
}

// Waits for any button to be pressed.
func WaitAny() Kind {
	for {
		bSync.Lock()

		f, _ := os.OpenFile(findFilename(), os.O_RDONLY, 0644)
		b := make([]byte, 16)
		f.Read(b)
		f.Close()

		codeData := b[10:11]
		valueData := b[12:13]

		buf := bytes.NewBuffer(codeData)
		var code int8
		binary.Read(buf, binary.BigEndian, &code)

		buf2 := bytes.NewBuffer(valueData)
		var value int8
		binary.Read(buf2, binary.BigEndian, &value)

		bSync.Unlock()

		if value == 0 {
			return Kind(code)
		}
	}
}
