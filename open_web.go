package main

import (
	"os/exec"
	"runtime"
)

func OpenUrl(url string) error {
	var exe string
	var args []string
	switch runtime.GOOS {
	case "windows":
		exe = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		exe = "open"
	default: // linux, freebsd, openbsd...
		exe = "xdg-open"
	}
	args = append(args, url)
	cmd := exec.Command(exe, args...)
	return cmd.Run()
}
