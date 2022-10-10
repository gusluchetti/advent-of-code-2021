package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

func main() {
	// FIXME: actually validate parse input is working
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

	counter := GetDepthIncreases(depthArray)
	fmt.Println("Part 1 Solution (Number of times depth increased): " + strconv.Itoa(counter))

	sectionsArray := GetSections(depthArray)
	counter = GetDepthIncreases(sectionsArray)
	fmt.Println("Part 2 Solution (Number of times depth increased in sections of 3): " + strconv.Itoa(counter))
}
