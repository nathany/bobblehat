package stick

import (
	"os"
	"strings"
	"syscall"
	"unsafe"
)

const (
	Enter = 28
	Up    = 103
	Left  = 105
	Right = 106
	Down  = 108
)

type Device struct {
	f      *os.File
	Events chan Event
}

func Open(name string) (*Device, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	d := &Device{f, make(chan Event, 4)}

	go d.pollEvents()

	return d, nil
}

func (d *Device) Name() string {
	var str [256]byte

	ioctl(d.f.Fd(), _EVIOCGNAME(256), uintptr(unsafe.Pointer(&str[0])))

	return strings.TrimRight(string(str[:]), "\x00")
}

func (d *Device) pollEvents() {
	defer close(d.Events)

	var e Event

	size := int(unsafe.Sizeof(e))
	buf := make([]byte, size*2)

	for {
		n, err := d.f.Read(buf)
		if err != nil {
			return
		}

		events := (*(*[1<<27 - 1]Event)(unsafe.Pointer(&buf[0])))[:n/size]

		for i := range events {
			if e := events[i]; e.Type == 0x01 && e.Value != 0 {
				d.Events <- e
			}
		}
	}
}

// Event represents a generic input event.
type Event struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

func ioctl(fd, name, v uintptr) error {
	_, _, errno := syscall.RawSyscall(syscall.SYS_IOCTL, fd, name, v)
	if errno == 0 {
		return nil
	}

	return errno
}

func _EVIOCGNAME(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x06, len)
}

const (
	_IOC_WRITE     = 0x1
	_IOC_READ      = 0x2
	_IOC_NRBITS    = 8
	_IOC_TYPEBITS  = 8
	_IOC_SIZEBITS  = 14
	_IOC_NRSHIFT   = 0
	_IOC_TYPESHIFT = _IOC_NRSHIFT + _IOC_NRBITS
	_IOC_SIZESHIFT = _IOC_TYPESHIFT + _IOC_TYPEBITS
	_IOC_DIRSHIFT  = _IOC_SIZESHIFT + _IOC_SIZEBITS
)

func _IOC(dir, t, nr, size int) uintptr {
	return uintptr((dir << _IOC_DIRSHIFT) | (t << _IOC_TYPESHIFT) |
		(nr << _IOC_NRSHIFT) | (size << _IOC_SIZESHIFT))
}
