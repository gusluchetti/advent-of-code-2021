package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gusluchetti/advent-of-code/utils"
)

func closestToAll(crabs []int) int {
	// median is closest to all
	sort.Ints(crabs)
	fmt.Printf("\nSorted crabs: %v", crabs)
	midIndex := (len(crabs)-1)/2
	fuel := 0
	for _, c := range crabs {
		midCrab := crabs[midIndex]
		// fmt.Printf("\nMid crab is %d", midIndex)
		fuel += int(math.Abs(float64(midCrab - c)))
		fmt.Printf("\nCost for %d to go to %d: %d", c, midCrab, fuel)
	}
	fmt.Printf("\nTotal fuel cost: %d", fuel)

	return fuel
}

func main() {
	file, err := os.Open("Y2021D07_input.txt")
	utils.Check(err)

	scanner := bufio.NewScanner(file)
	var tokens []string
	// effectively a list of lines (strings)
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}
	// first element of tokens is literally whole input
	split := strings.Split(tokens[0], ",")
	var crabs []int
	for _, s := range split {
		num, err := strconv.Atoi(s)
		utils.Check(err)
		crabs = append(crabs, num)
	}

	testCrabs := []int{16,1,2,0,4,2,7,1,2,14}
	leastFuel := closestToAll(testCrabs)
	if leastFuel == 37 {
		fmt.Println("\nOK!")

		leastFuel = closestToAll(crabs)
		fmt.Println(leastFuel)
	}
}
