// +build darwin

package main

import (
	"os/exec"
)

func notifier(title string, message string) error {
	cmd := exec.Command("terminal-notifier", "-title", title, "-message", message)
	return cmd.Run()
}
