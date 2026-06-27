package example

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"default", "world", "Hello, world!"},
		{"custom", "Go", "Hello, Go!"},
		{"empty", "", "Hello, !"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet(tt.input)
			if got != tt.expected {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}