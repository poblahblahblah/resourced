package wsclient

import (
	"strings"
	"testing"
)

func TestNewClien(t *testing.T) {
	_, _, err := NewClient("http://localhost:55555", "http://localhost:55655/api/ws", nil)
	if err != nil && !strings.Contains(err.Error(), "connection refused") {
		t.Fatalf("Creating new client with basic settings should not fail. Error: %v", err)
	}
}
