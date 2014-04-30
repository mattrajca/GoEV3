// PRIVATE helper functions
package utilities

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

var gLocks map[string]*sync.RWMutex
var gCreationLock = &sync.Mutex{}

func ensureLockForFilename(filename string) {
	gCreationLock.Lock()

	if _, ok := gLocks[filename]; !ok {
		if gLocks == nil {
			gLocks = make(map[string]*sync.RWMutex)
		}

		gLocks[filename] = &sync.RWMutex{}
	}

	gCreationLock.Unlock()
}

func ReadStringValue(filename string, basename string) string {
	actualFilename := fmt.Sprintf("%s/%s", filename, basename)
	ensureLockForFilename(actualFilename)

	gLocks[actualFilename].RLock()

	data, _ := ioutil.ReadFile(actualFilename)
	str := string(data)

	gLocks[actualFilename].RUnlock()

	return strings.TrimSpace(str)
}

func ReadIntValue(filename string, basename string) int64 {
	str := ReadStringValue(filename, basename)
	result, _ := strconv.ParseInt(str, 10, 16)

	return result
}

func ReadUInt8Value(filename string, basename string) uint8 {
	return uint8(ReadIntValue(filename, basename))
}

func ReadUInt16Value(filename string, basename string) uint16 {
	return uint16(ReadIntValue(filename, basename))
}

func ReadInt16Value(filename string, basename string) int16 {
	return int16(ReadIntValue(filename, basename))
}

func ReadUInt32Value(filename string, basename string) uint32 {
	return uint32(ReadIntValue(filename, basename))
}

func WriteStringValue(filename string, basename string, value string) {
	actualFilename := fmt.Sprintf("%s/%s", filename, basename)
	ensureLockForFilename(actualFilename)

	gLocks[actualFilename].Lock()

	data := []byte(value)

	ioutil.WriteFile(actualFilename, data, 0644)

	gLocks[actualFilename].Unlock()
}

func WriteIntValue(filename string, basename string, value int64) {
	WriteStringValue(filename, basename, strconv.FormatInt(value, 10))
}

func WriteUIntValue(filename string, basename string, value uint64) {
	WriteStringValue(filename, basename, strconv.FormatUint(value, 10))
}
