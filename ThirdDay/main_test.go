package main

import (
	"testing"
)

func TestPowerConsumption (t * testing.T) {
	testArray := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	result := GetPowerConsumption(testArray)
	if result != 198 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", result, 198)
	}
}