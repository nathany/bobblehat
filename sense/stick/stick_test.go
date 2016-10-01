package stick

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestStickWithEmptyDeviceFile(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	for i := 0; i < 100; i++ {
		tmpfile.Write([]byte{0x30, 0xab, 0xe6, 0x57, 0xc0, 0x80, 0x8, 0x0, 0x1, 0x0, 0x1c, 0x0, 0x1, 0x0, 0x0, 0x0, 0x30, 0xab, 0xe6, 0x57, 0xc0, 0x80, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
	}
	tmpfile.Sync()

	input, err := Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if got, want := input.Name(), ""; got != want {
		t.Fatalf("input.Name() = %q, want %q", got, want)
	}

	timeout := time.After(5 * time.Second)

	for {
		select {
		case e := <-input.Events:
			if e.Type == 1 && e.Code == Enter {
				t.Logf("Enter was pressed!")
				return
			}
		case <-timeout:
			t.Fatal("could not read any events")
		}
	}
}
