package main

import (
	"bytes"
	"os"
	"testing"
)

// TestMainFunction tests the main function to ensure it handles NTP errors properly
func TestMainFunction(t *testing.T) {
	// Save the original stderr and stdout
	originalStderr := os.Stderr
	originalStdout := os.Stdout

	// Create a pipe to capture stderr
	rStderr, wStderr, _ := os.Pipe()
	os.Stderr = wStderr

	// Create a pipe to capture stdout
	rStdout, wStdout, _ := os.Pipe()
	os.Stdout = wStdout

	// Run the main function
	main()

	// Close the pipes
	wStderr.Close()
	wStdout.Close()

	// Read captured output
	var stderrBuf bytes.Buffer
	var stdoutBuf bytes.Buffer
	stderrBuf.ReadFrom(rStderr)
	stdoutBuf.ReadFrom(rStdout)

	// Restore original stderr and stdout
	os.Stderr = originalStderr
	os.Stdout = originalStdout

	// Check if stderr contains any output (indicating an error)
	if stderrBuf.Len() != 0 {
		t.Errorf("Expected no errors, but got: %s", stderrBuf.String())
	}

	// Check if stdout contains the expected output
	if stdoutBuf.Len() == 0 {
		t.Error("Expected output to stdout, but got none")
	}
}
