package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/gusluchetti/advent-of-code/utils"
)

func closestToAll(crabs []int) int {
	lastIndex := len(crabs)-1
	initPoint := crabs[rand.Intn(lastIndex + 1)]

	var bestCost float64 = float64(math.MaxUint)
	var cost float64 = 0
	for i:=0;i<lastIndex;i++ {
		distance := float64(initPoint - crabs[i])
		cost += math.Abs(distance)
	}
	if cost < bestCost {
		bestCost = cost
	}

	return int(bestCost)
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
		fmt.Println("OK")
	}
}
