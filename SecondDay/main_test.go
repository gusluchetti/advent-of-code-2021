package main

import (
	"testing"
)

func TestCommands (t * testing.T) {
	testArray := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	result := ReturnMultipliedCommands(testArray)
	if result != 150 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", result, 150)
	}
}