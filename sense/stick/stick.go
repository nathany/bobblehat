package stick

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Joystick events
var (
	Right = 0
	Left  = 1
	Up    = 2
	Down  = 3
	Enter = 4
)

// the joy stick
var joyStickDevice string

func init() {
	var err error
	joyStickDevice, err = getDevice("Raspberry Pi Sense HAT Joystick")
	if err != nil {
		panic(err)
	}
}

// ReadEvent from the joystick.
func ReadEvent() int {
	file, err := os.Open(joyStickDevice)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes, %q\n", count, data[:count])

	return count
}

func getDevice(name string) (string, error) {
	matches, err := filepath.Glob("/sys/class/input/event*")
	if err != nil {
		return "", err
	}

	for _, dir := range matches {
		b, err := ioutil.ReadFile(filepath.Join(dir, "device/name"))
		if err != nil {
			continue
		}
		fbName := strings.TrimSpace(string(b))
		if fbName == name {
			dev := filepath.Join("/dev/input", filepath.Base(dir))

			return dev, nil
		}
	}
	return "", errStickDeviceNotFound
}

// errors
var (
	errStickDeviceNotFound = errors.New("stick device not found")
)
