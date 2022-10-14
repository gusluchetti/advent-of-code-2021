package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gusluchetti/advent-of-code/utils"
)

func removeCandidate(curIndex int, fishMatrix [][]int) [][]int{
	fmt.Printf("\nFirst item %v was deleted\n", fishMatrix[curIndex])
	fishMatrix = fishMatrix[1:]

	return fishMatrix
}

func SmartFishSim(fish []int, days int, firstSpawnAge int, respawnAge int) int {
	fishMatrix := [][]int{}
	for _, f := range fish {
		fishy := []int{f, days}
		fishMatrix = append(fishMatrix, fishy)
	}

	sum := 0
	for len(fishMatrix) > 0 {
		fmt.Printf("\nCurrent Sum: %d", sum)
		fmt.Printf("\nStart: %v", fishMatrix)
		i := 0
		// generate items that this item generates
		sub := fishMatrix[i][0] + 1
		remainder := fishMatrix[i][1]-sub
		rr := remainder - (respawnAge+1)
		for rr >= -1 {
			new := []int{firstSpawnAge, rr}
			fishMatrix = append(fishMatrix, new)
			fmt.Printf("\nGenerated new candidate %v", new)
			rr = rr - (respawnAge+1)
		}
		// item is removed when it can no longer generate
		fishMatrix = removeCandidate(i, fishMatrix)
		sum += 1
	}
	return sum
}

func DumbFishSim(fish []int, days int) int {
	fmt.Printf("\nDay 0: %v", fish)
	for d:=1;d<=days;d++ {
		for i:=0;i<len(fish);i++ {
			fish[i] = fish[i] - 1
			if fish[i] == -1 {
				fish[i] = 6
				fish = append(fish, 9)
			}
		}
		fmt.Printf("\nDay %d: %v, len=%d", d, fish, len(fish))
	}
	return len(fish)
}

func main() {
	file, err := os.Open("Y2021D06_input.txt")
	utils.Check(err)

	scanner := bufio.NewScanner(file)
	var tokens []string
	// effectively a list of lines (strings)
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}
	// first element of tokens is literally whole input
	split := strings.Split(tokens[0], ",")
	var fish []int
	for _, s := range split {
		num, err := strconv.Atoi(s)
		utils.Check(err)
		fish = append(fish, num)
	}

	firstSpawnAge := 8
	respawnAge := 6
	// _ = DumbFishSim([]int{3, 2}, days) //3,2, 20 days, 14
	sum := SmartFishSim([]int{3, 4, 3, 1, 2}, 18, firstSpawnAge, respawnAge)
	fmt.Printf("\n\nTotal fish population is %d", sum)
}
