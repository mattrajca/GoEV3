// Provides APIs for interacting with EV3's buttons.
package Button

import (
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"sync"
)

var bPressedMap map[Kind]bool

var bSync = &sync.Mutex{}
var bMapLock = &sync.Mutex{}

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

		bSync.Unlock()

		codeData := b[10:11]
		valueData := b[12:13]

		buf := bytes.NewBuffer(codeData)
		var code int8
		binary.Read(buf, binary.BigEndian, &code)

		buf2 := bytes.NewBuffer(valueData)
		var value int8
		binary.Read(buf2, binary.BigEndian, &value)

		if value == 0 && code == int8(kind) {
			break
		}
	}
}

// Watches for button presses in a background thread. Use `IsPressed` to query the current state.
func Watch() {
	bPressedMap = make(map[Kind]bool)

	go func() {
		for {
			bSync.Lock()

			f, _ := os.OpenFile(findFilename(), os.O_RDONLY, 0644)
			b := make([]byte, 16)
			f.Read(b)
			f.Close()

			bSync.Unlock()

			codeData := b[10:11]
			valueData := b[12:13]

			buf := bytes.NewBuffer(codeData)
			var code int8
			binary.Read(buf, binary.BigEndian, &code)

			buf2 := bytes.NewBuffer(valueData)
			var value int8
			binary.Read(buf2, binary.BigEndian, &value)

			bMapLock.Lock()

			if value == 0 {
				bPressedMap[Kind(code)] = true
			} else {
				bPressedMap[Kind(code)] = false
			}

			bMapLock.Unlock()
		}
	}()
}

// Checks if the given button is currently pressed. This call must be preceded with a call to `Watch`.
func IsPressed(kind Kind) bool {
	var result bool

	bMapLock.Lock()
	result = bPressedMap[kind]
	bMapLock.Unlock()

	return result
}

// Waits for any button to be pressed.
func WaitAny() Kind {
	for {
		bSync.Lock()

		f, _ := os.OpenFile(findFilename(), os.O_RDONLY, 0644)
		b := make([]byte, 16)
		f.Read(b)
		f.Close()

		bSync.Unlock()

		codeData := b[10:11]
		valueData := b[12:13]

		buf := bytes.NewBuffer(codeData)
		var code int8
		binary.Read(buf, binary.BigEndian, &code)

		buf2 := bytes.NewBuffer(valueData)
		var value int8
		binary.Read(buf2, binary.BigEndian, &value)

		if value == 0 {
			return Kind(code)
		}
	}
}
