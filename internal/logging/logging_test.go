package logging

import (
	"testing"
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.Out = &buf
	logger.Formatter = new(logrus.TextFormatter)
	log = logger

	cases := []struct {
		input    string
		expected logrus.Level
	}{
		{"debug", logrus.DebugLevel},
		{"info", logrus.InfoLevel},
		{"warning", logrus.WarnLevel},
		{"error", logrus.ErrorLevel},
		{"fatal", logrus.FatalLevel},
		{"nonexistent", logrus.InfoLevel}, // default level
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			Level(tc.input)
			assert.Equal(t, tc.expected, log.Level)
		})
	}
}

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	logger := logrus.New()
	logger.Out = &buf
	logger.Formatter = new(logrus.TextFormatter)
	logger.Level = logrus.DebugLevel  // Add this line to set the log level to DebugLevel
	log = logger

	msg := "debug message"
	Debug(msg)
	assert.Contains(t, buf.String(), msg)
}