package psn

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	// Test case 1: Valid options
	opts1 := &Options{
		Lang:   "en",
		Region: "us",
	}
	client1, err1 := NewClient(opts1)
	if err1 != nil {
		t.Errorf("Expected no error for valid options, but got: %v", err1)
	}
	if client1 == nil {
		t.Error("Expected a non-nil client for valid options, but got nil")
	}

	// Test case 2: Invalid language
	opts2 := &Options{
		Lang:   "invalid-lang",
		Region: "us",
	}
	_, err2 := NewClient(opts2)
	if err2 == nil {
		t.Error("Expected an error for invalid language, but got nil")
	}

	// Test case 3: Invalid region
	opts3 := &Options{
		Lang:   "en",
		Region: "invalid-region",
	}
	_, err3 := NewClient(opts3)
	if err3 == nil {
		t.Error("Expected an error for invalid region, but got nil")
	}
}