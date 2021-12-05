package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	
	var depthArray []int
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		depthArray = append(depthArray, depth)
	}

	counter := GetDepthIncreases(depthArray)
	fmt.Println("\n Number of times depth increased: " + strconv.Itoa(counter))
}
