package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func PrintDiagram(diagram [][]int) {
	for _, line := range diagram {
		fmt.Println(line)
	}
	fmt.Println(" --- ")
}

func NumberOfOverlappingPoints(diagram [][]int) int {
	overlaps := 0
	for _, line := range diagram {
		for _, number := range line {
			if number > 1 {
				overlaps++
			}
		}
	}
	return overlaps
}

func AddInputToDiagram(tokens []string, diagram [][]int) [][]int {
	for _, token := range tokens {
		res := regexp.MustCompile(`\d+`)
		n := res.FindAllString(token, -1)
		var start int
		var end int
		// numbers[0], [1] ->  initial coordinates
		// numbers[2], [3] -> final coordinates
		if n[1] == n[3] { // horizontal line
			baseY, _ := strconv.Atoi(n[1])
			first, _ := strconv.Atoi(n[0])
			second, _ := strconv.Atoi(n[2])
			if first > second {
				start = second
				end = first
			} else {
				start = first
				end = second
			}
			for x := start; x <= end; x++ {
				diagram[baseY][x]++
			}
		} else if n[0] == n[2] { // vertical line
			baseX, _ := strconv.Atoi(n[0])
			first, _ := strconv.Atoi(n[1])
			second, _ := strconv.Atoi(n[3])
			if first > second {
				start = second
				end = first
			} else {
				start = first
				end = second
			}

			for y := start; y <= end; y++ {
				diagram[y][baseX]++
			}
		} else {
			fmt.Println("Diagonal lines are not implemented yet!")
		}

		// PrintDiagram(diagram)
	}

	return diagram
}

func BuildDynamicDiagram(size int) [][]int {
	diagram := make([][]int, size)
	for i := 0; i < size; i++ {
		diagram[i] = make([]int, size)
	}

	return diagram
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var tokens []string
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	diagram := BuildDynamicDiagram(1000)
	diagram = AddInputToDiagram(tokens, diagram)
	overlaps := NumberOfOverlappingPoints(diagram)
	fmt.Printf("\nTwo lines overlap on %d points.\n", overlaps)
}
