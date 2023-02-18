package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeConversion(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "testSucceed",
			input:    "03:04:05PM",
			expected: "15:04:05",
		},
		{
			name:     "testSucceed",
			input:    "12:04:05PM",
			expected: "12:04:05",
		},
		{
			name:     "testSucceed",
			input:    "03:04:05AM",
			expected: "03:04:05",
		},
		{
			name:     "testSucceed",
			input:    "12:04:05AM",
			expected: "00:04:05",
		},
		{
			name:     "testErrorAM",
			input:    "gh:04:05AM",
			expected: `strconv.Atoi: parsing "gh": invalid syntax`,
		},
		{
			name:     "testErrorPM",
			input:    "gh:04:05PM",
			expected: `strconv.Atoi: parsing "gh": invalid syntax`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get := timeConversion(tt.input)
			log.Print(get)
			assert.Equal(t, tt.expected, get)

		})
	}

}
