package main

import (
	"fmt"
	"os"
	"os/exec"
)

// checkDependencies verifies that required external tools are installed
func checkDependencies() error {
	// Check for mpv
	if _, err := exec.LookPath("mpv"); err != nil {
		return fmt.Errorf("mpv not found. Please install it: https://mpv.io")
	}

	return nil
}

// Play streams audio or video from the given URL
func Play(url string, videoMode bool) error {
	if err := checkDependencies(); err != nil {
		return err
	}

	var args []string

	if videoMode {
		// Stream both video and audio
		args = []string{url}
	} else {
		// Stream audio only (no video)
		args = []string{"--no-video", url}
	}

	// Run mpv with the URL
	cmd := exec.Command("mpv", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Printf("Playing %s (video: %v)...\n", url, videoMode)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to play stream: %w", err)
	}

	return nil
}
