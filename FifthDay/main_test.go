package main

import (
	"fmt"
	"testing"
)

func TestVents(t *testing.T) {
	tokens := []string{"0,9 -> 5,9", "8,0 -> 0,8", "9,4 -> 3,4", "2,2 -> 2,1", "7,0 -> 7,4",
		"6,4 -> 2,0", "0,9 -> 2,9", "3,4 -> 1,4", "0,0 -> 8,8", "5,5 -> 8,2"}

	diagram := BuildDynamicDiagram(10)
	diagram = AddInputToDiagram(tokens, diagram)
	overlaps := NumberOfOverlappingPoints(diagram)
	fmt.Printf("\nTwo lines overlap on %d points.\n", overlaps)

	if overlaps != 5 {
		t.Errorf("Result is incorrect, expected %d, got %d", 5, overlaps)
	}
}
