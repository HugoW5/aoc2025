package main

import (
	"fmt"
	"testing"
)

func TestCheckInvalidID(t *testing.T) {
	tests := []struct {
		isPattern bool
		input     string
	}{
		{isPattern: true, input: "123123"},
		{isPattern: true, input: "11111"},
		{isPattern: true, input: "22"},
		{isPattern: false, input: "45678"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("input: %s", tt.input), func(t *testing.T) {
			isPattern := checkForPattern(tt.input)

			if isPattern != tt.isPattern {
				t.Error("Not correct")
			}
		})
	}
}
