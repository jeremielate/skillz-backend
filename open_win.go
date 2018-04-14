// +build windows

package main

import "os/exec"

func OpenUrl(url string) error {
	cmd := exec.Command("start", url)
	return cmd.Run()
}
