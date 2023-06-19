package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/gusluchetti/advent-of-code/utils"
)

func GetDepthIncreases(depthArray []int) int {
	var counter int = 0;
		for i:=1; i < len(depthArray); i++ {
			if depthArray[i] > depthArray[i-1] {
				counter++
			}
		}
		return counter
}

func GetSections(depthArray []int) []int {
	maxLength := len(depthArray) - 2
	var sectionsArray []int
		for i:=0; i < maxLength; i++ {
			sectionsArray = append(sectionsArray, (depthArray[i] + depthArray[i+1] + depthArray[i+2]))
		}

		return sectionsArray
}

func setup() []int {
	path, err := os.Getwd()
	utils.Check(err)
	tokens := utils.ParseInput(path+"/src/Y2021/D01/Y2021D01_input.txt")

	var depthArray []int
	for _, token := range tokens {
		depth, err := strconv.Atoi(token)
		if err != nil {
			log.Fatal(err)
		}
		depthArray = append(depthArray, depth)
	}

	return depthArray
}

func PartOne(depthArray []int) int {
	counter := GetDepthIncreases(depthArray)
	return counter
}

func TestOne (t * testing.T) {
	testArray := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	counter := PartOne(testArray)
	if counter != 7 {
		t.Errorf("Counter was incorrect, got: %d, expected: %d", counter, 7)
	}
}

func PartTwo(depthArray []int) int {
	sectionsArray := GetSections(depthArray)
	counter := GetDepthIncreases(sectionsArray)
	return counter
}

func TestTwo (t * testing.T) {
	testArray := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	counter := PartTwo(testArray)
	if counter != 5 {
		t.Errorf("Counter was incorrect, got: %d, expected: %d", counter, 5)
	}
}

func main() {
	depthArray := setup()
	partOne := PartOne(depthArray)
	fmt.Println(partOne)
	partTwo := PartTwo(depthArray)
	fmt.Println(partTwo)
}
