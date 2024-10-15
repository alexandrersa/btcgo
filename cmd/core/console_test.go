package core

import (
	"testing"
)

func TestReadCPusForUse(t *testing.T) {
	// Define test cases
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "Valid CPUS number",
			input: 4,
			want:  4,
		},
		{
			name:  "Invalid CPUS number",
			input: 55,
			want:  0,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readCPUsForUse(); got != tt.want {
				t.Errorf("readCPUsForUse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromptRangeNumber(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "Valid range number",
			input: 2,
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := promptRangeNumber(); got != tt.want {
				t.Errorf("promptRangeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromptMods(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Mode 1",
			want: 1,
		},
		{
			name: "Mode 2",
			want: 2,
		},
		{
			name: "Mode 3",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := promptMods(3); got != tt.want {
				t.Errorf("promptMods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestData(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "RequestData test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RequestData()
		})
	}
}
