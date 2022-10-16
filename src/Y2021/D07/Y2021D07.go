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

func PartOne(crabs []int) int {
	// median is closest to all
	sort.Ints(crabs)
	bestIndex := (len(crabs)-1)/2
	fmt.Printf("\nMid index is %d", bestIndex)
	fuel := 0
	for _, c := range crabs {
		position := crabs[bestIndex]
		// fmt.Printf("\nMid crab is %d", midIndex)
		fuel += int(math.Abs(float64(position - c)))
		// fmt.Printf("\nCost for %d to go to %d: %d", c, midCrab, fuel)
	}
	fmt.Printf("\nTotal fuel cost: %d", fuel)

	return fuel
}

func PartTwo(crabs []int) int {
	// given fuel is not constant cost
	sort.Ints(crabs)
	bestIndex := -1
	leastFuel := math.MaxInt

	for i:=0;i<crabs[len(crabs)-1];i++ {
		fmt.Printf("\nTesting index %d", i)
		fuel := 0
		for _, curCrab := range crabs {
			distance := math.Abs(float64(i-curCrab))
			fuel += int(distance*(1+distance)/2)
			fmt.Printf("\nTotal fuel cost (%d -> %d): %d", curCrab, i, fuel)
		}
		if fuel < leastFuel {
			leastFuel = fuel
			bestIndex = i
		}
		fmt.Printf("\nCur. best position is %d", crabs[bestIndex])
	}

	fmt.Printf("\nTotal fuel cost: %d", leastFuel)
	return leastFuel
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
	fmt.Printf("\n\nPart One Test Solution")
	leastFuel := PartOne(testCrabs)
	if leastFuel == 37 {
	fmt.Printf("\n\nPart One REAL Solution")
		leastFuel = PartOne(crabs)
	}

	fmt.Printf("\n\nPart Two Test Solution")
	leastFuel = PartTwo(testCrabs)
	if leastFuel == 168 {
	fmt.Printf("\n\nPart Two Real Solution")
		leastFuel = PartTwo(crabs)
	}
}

// from what i've gathered online
// median for constant fuel makes sense, literal mid point of all points
// mean for ever-increasing fuel didn't make sense, if I could plot out the graph
// of fuel cost, that should make more sense (graph is concave up)
// how does binary search come into this?
//
// reddit comment - the part 1 destination is just the median,
// and the part 2 destination is close to the mean.
// Can you use math to show that this is true in general?
