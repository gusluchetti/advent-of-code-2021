package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	tokens := ReadInputAsTokens("input.txt")
	var depthArray []int
	for _, token := range tokens {
		depth, err := strconv.Atoi(token)
		if err != nil {
			log.Fatal(err)
		}
		depthArray = append(depthArray, depth)
	}

	counter := GetDepthIncreases(depthArray)
	fmt.Println("\n Number of times depth increased: " + strconv.Itoa(counter))

	sectionsArray := GetSections(depthArray)
	counter = GetDepthIncreases(sectionsArray)

	fmt.Println("\n Number of times depth increased (on sections): " + strconv.Itoa(counter))
}
