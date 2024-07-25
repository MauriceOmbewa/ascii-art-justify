package terminal

import (
    "os"
    "syscall"
    "unsafe"
)

// GetTerminalSize returns the width and height of the terminal.
func GetTerminalSize() (width, height int, err error) {
	var winsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}

	// Get the file descriptor for stdout
	fd := int(os.Stdout.Fd())

	// Retrieve the terminal size
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&winsize)))
	if errno != 0 {
		return 0, 0, errno
	}

	return int(winsize.Col), int(winsize.Row), nil
}