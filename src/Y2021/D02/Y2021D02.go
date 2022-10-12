package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/gusluchetti/advent-of-code/utils"
)

func ReturnCommands(tokens []string) int{
	commands := make([]int, 2)
	for _, token := range tokens {
		split := strings.Fields(token)
		value, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		switch split[0] {
		case "forward":
			commands[0] += value
		case "up":
			commands[1] -= value
		case "down":
			commands[1] += value
		}
	}

	return commands[0] * commands[1]
}

func ReturnCommandsWithAim(tokens []string) int {
	commands := make([]int, 2)
	aim := 0

	for _, token := range tokens {
		split := strings.Fields(token)
		value, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		switch split[0] {
		case "forward":
			commands[0] += value
			commands[1] += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

	return commands[0] * commands[1]
}

func setup() []string {
	path, err := os.Getwd()
	utils.Check(err)
	return utils.ParseInput(path+"/src/Y2021/D02/Y2021D02_input.txt")
}

func PartOne(tokens []string) int {
	return ReturnCommands(tokens)
}

func TestOne (t * testing.T) {
	testArray := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	result := PartOne(testArray)
	if result != 150 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", result, 150)
	}
}

func PartTwo(tokens []string) int {
	return ReturnCommandsWithAim(tokens)
}

func TestCommandsWithAim (t * testing.T) {
	testArray := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	result := PartTwo(testArray)
	if result != 900 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", result, 900)
	}
}

func main() {
	tokens := setup()
	partOne := PartOne(tokens)
	fmt.Println(partOne)
	partTwo := PartTwo(tokens)
	fmt.Println(partTwo)
}
