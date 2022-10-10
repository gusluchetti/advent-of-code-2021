package main

import (
	"fmt"
	"testing"
)

func TestDepth (t * testing.T) {
	testArray := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	counter := GetDepthIncreases(testArray)
	if counter != 7 {
		t.Errorf("Counter was incorrect, got: %d, expected: %d", counter, 7)
	}
}

func TestSectionsIncreases (t * testing.T) {
	testArray := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	sectionsArray := GetSections(testArray)
	counter := GetDepthIncreases(sectionsArray)
	fmt.Println(sectionsArray)
	fmt.Println(counter)

	if counter != 5 {
		t.Errorf("Counter was incorrect, got: %d, expected: %d", counter, 5)
	}
}
