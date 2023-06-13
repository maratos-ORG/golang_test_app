package config

import (
	"testing"
)

func TestBackendParameters(t *testing.T) {
	// Test instance of BackendParameters
	params := BackendParameters{
		DBURL:       stringPointer("localhost"),
		Port:        stringPointer("8080"),
		LogLevel:    stringPointer("INFO"),
		ShowVersion: boolPointer(true),
	}

	if *params.DBURL != "localhost" {
		t.Errorf("DBURL was incorrect, got: %s, want: %s", *params.DBURL, "localhost")
	}

	if *params.Port != "8080" {
		t.Errorf("Port was incorrect, got: %s, want: %s", *params.Port, "8080")
	}

	if *params.LogLevel != "INFO" {
		t.Errorf("LogLevel was incorrect, got: %s, want: %s", *params.LogLevel, "INFO")
	}

	if *params.ShowVersion != true {
		t.Errorf("ShowVersion was incorrect, got: %v, want: %v", *params.ShowVersion, true)
	}
}

// Helpers for creating pointer to string and bool
func stringPointer(s string) *string {
	return &s
}

func boolPointer(b bool) *bool {
	return &b
}

// explanation https://chat.openai.com/share/a9a75726-a9f8-4053-b884-d58e285677fd