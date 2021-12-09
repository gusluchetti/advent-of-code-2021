package Tests

import (
	"testing"
)

func TestDepth (t * testing.T) {
	testArray := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	counter := GetDepthIncreases(testArray)
	if counter != 7 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", counter, 7)
	}
}