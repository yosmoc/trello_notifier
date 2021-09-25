// +build linux

package main

import (
	"os/exec"
)

func notifier(title string, message string) error {
	cmd := exec.Command("notify-send", title, message)
	return cmd.Run()
}
