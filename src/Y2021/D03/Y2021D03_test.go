package Y2021D03

import (
	"testing"
	"strings"
)

func TestPowerConsumption (t * testing.T) {
	tokens := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	bitSize := len(strings.Split(tokens[1], ""))
	mostCommonBitsPerColumn := GetMostCommonBitsPerColumn(make([]int, bitSize), tokens)
	result := GetPowerConsumption(tokens, mostCommonBitsPerColumn, bitSize)
	
	if result != 198 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", result, 198)
	}
}

func TestLifeSupportRating (t * testing.T) {
	tokens := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	result := GetLifeRating(tokens)
	
	if result != 230 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", result, 230)
	}
}
