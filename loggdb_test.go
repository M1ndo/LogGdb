package loggdb_test

import (
	"bytes"
	"testing"

	"github.com/m1ndo/LogGdb"
)

func TestLogger(t *testing.T) {
	// Create a temporary directory for the log file
	// Create a logger instance

	logger := &loggdb.Logger{
		LogDir: "logs",
	}
	err := logger.NewLogger()
	if err != nil {
		panic(err)
	}

	// Set up a buffer to capture the log output
	var buf bytes.Buffer
	logger.SetOutput(&buf)

	// Write some log messages
	logger.Info("This is an info message")
	logger.Error("This is an error message")

	// Get the captured log output from the buffer
	logOutput := buf.String()

	// Check if the log output contains the expected messages
	if !contains(logOutput, "This is an info message") {
		t.Errorf("Expected log output to contain 'This is an info message', got: %s", logOutput)
	}
	if !contains(logOutput, "This is an error message") {
		t.Errorf("Expected log output to contain 'This is an error message', got: %s", logOutput)
	}
}

// Helper function to check if a string contains a substring
func contains(str, substr string) bool {
	return bytes.Contains([]byte(str), []byte(substr))
}
