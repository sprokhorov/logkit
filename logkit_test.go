package logkit

import (
	"bytes"
	"log"
	"testing"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0) // Disable timestamps
	defer func() {
		log.SetOutput(nil)          // Reset log output
		log.SetFlags(log.LstdFlags) // Restore the default flags
	}()

	f()
	return buf.String()
}

func TestDefaultLogger_Debug(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureOutput(func() {
		logger.Debug("debug message")
	})
	expected := "debug message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestDefaultLogger_Debugf(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureOutput(func() {
		logger.Debugf("debug %s", "message")
	})
	expected := "debug message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

// Similar adjustments for other test cases...

func TestDefaultLogger_Info(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureOutput(func() {
		logger.Info("info message")
	})
	expected := "info message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestDefaultLogger_Infof(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureOutput(func() {
		logger.Infof("info %s", "message")
	})
	expected := "info message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestDefaultLogger_Warn(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureOutput(func() {
		logger.Warn("warn message")
	})
	expected := "warn message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestDefaultLogger_Warnf(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureOutput(func() {
		logger.Warnf("warn %s", "message")
	})
	expected := "warn message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestDefaultLogger_Error(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureOutput(func() {
		logger.Error("error message")
	})
	expected := "error message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestDefaultLogger_Errorf(t *testing.T) {
	logger := &DefaultLogger{}
	output := captureOutput(func() {
		logger.Errorf("error %s", "message")
	})
	expected := "error message\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}
