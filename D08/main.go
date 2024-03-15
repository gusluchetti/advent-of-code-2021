package main

import (
	"bufio"
	"fmt"
	"os"
)

func PartOne(lines []int) int {
	return -1
}

func PartTwo(lines []int) int {
	return -1
}

func main() {
	file, err := os.Open("Y2021D08_input.txt")
	if err != nil {
		panic("oh no!")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Print(lines)
}
