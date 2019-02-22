package main

import (
	"log"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	RunServer()
}

func RunServer() {
	app := NewApp()
	var command string
	switch runtime.GOOS {
	case "windows":
		command = "explorer"
	case "darwin":
		command = "open"
	case "linux":
		command = "xdg-open"
	default:
		log.Fatalf("unknown operating system: %s", runtime.GOOS)
	}
	go func() {
		time.Sleep(3 * time.Second)
		cmd := exec.Command(command, "http://localhost:3001/")
		cmd.Run()
	}()

	app.Start(":3001")
}
