package Tests

import (
	"fmt"
	"testing"

	FirstDay "github.com/gusluchetti/advent-of-code-2021/FirstDay"
)

func TestSections (t * testing.T) {
	testArray := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	sectionsArray := FirstDay.GetSections(testArray)
	counter := FirstDay.GetDepthIncreases(sectionsArray)
	fmt.Println(sectionsArray)
	fmt.Println(counter)

	if counter != 5 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", counter, 5)
	}
}