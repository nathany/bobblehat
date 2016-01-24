package stick

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gophergala2016/bobblehat/sense/screen"
	"github.com/gophergala2016/bobblehat/sense/screen/color"
)

var (
	RIGHT = 0
	LEFT  = 1
	UP    = 2
	DOWN  = 3
	ENTER = 4
)

func Draw() {
	fb := screen.NewFrameBuffer()

	count := 0
	for {
		count = readEvent()
		fmt.Printf("Count is-->%d", count)

		rand.Seed(time.Now().UnixNano())
		direction := rand.Intn(5)

		switch direction {
		case RIGHT:
			fmt.Println("right")
			for i := 0; i < 8; i++ {
				for j := 4; j < 8; j++ {
					fb.SetPixel(i, j, color.Red)
				}
			}
			break
		case LEFT:
			fmt.Println("left")
			for i := 0; i < 8; i++ {
				for j := 0; j < 4; j++ {
					fb.SetPixel(i, j, color.Black)
				}
			}
			break
		case UP:
			fmt.Println("up")
			for i := 0; i < 4; i++ {
				for j := 0; j < 8; j++ {
					fb.SetPixel(i, j, color.Blue)
				}
			}
			break
		case DOWN:
			fmt.Println("down")
			for i := 4; i < 8; i++ {
				for j := 0; j < 8; j++ {
					fb.SetPixel(i, j, color.Green)
				}
			}
			break
		case ENTER:
			fmt.Println("reset")
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					fb.SetPixel(i, j, color.White)
				}
			}
			break
		default:
			fmt.Println("waiting to be pressed")
		}
		time.Sleep(10)
	}

	screen.Draw(fb)
}

func Clear() {
	screen.Clear()
}

// the joy stick
var joyStickDevice string
var stickEventDirection int

func init() {
	var err error
	joyStickDevice, err = getDevice("Raspberry Pi Sense HAT Joystick")
	if err != nil {
		panic(err)
	}
}

func readEvent() int {
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
