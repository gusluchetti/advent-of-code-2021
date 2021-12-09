package main

import (
	"fmt"
	"log"
	"strconv"

	Utils "github.com/gusluchetti/advent-of-code-2021/utils"
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
	tokens := Utils.ReadInputAsTokens("input.txt")
	var depthArray []int
	for _, token := range tokens {
		depth, err := strconv.Atoi(token)
		if err != nil {
			log.Fatal(err)
		}
		depthArray = append(depthArray, depth)
	}

	counter := GetDepthIncreases(depthArray)
	fmt.Println("Number of times depth increased: " + strconv.Itoa(counter))

	sectionsArray := GetSections(depthArray)
	counter = GetDepthIncreases(sectionsArray)

	fmt.Println("\n Number of times depth increased (on sections): " + strconv.Itoa(counter))
}
