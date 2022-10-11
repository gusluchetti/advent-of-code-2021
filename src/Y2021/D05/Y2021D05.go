package main

import (
	"bufio"
	"testing"

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

func OrganizeCoordinates(first int, second int) (int, int) {
	var start int
	var end int

	if first > second {
		start = second
		end = first
	} else {
		start = first
		end = second
	}

	return start, end
}

func GetInstructions(first int, second int) string {
	if first-second < 0 {
		return "+"
	} else if first-second > 0 {
		return "-"
	} else {
		return "="
	}
}

func ApplyInstructions(diagram [][]int, xInst string, yInst string, x1 int, x2 int, y1 int, y2 int) [][]int {
	fmt.Printf("\nInstructions found: %s for X and %s for Y [%d, %d -> %d, %d]", xInst, yInst, x1, y1, x2, y2)
	var start int
	var end int
	if (xInst == "-" || xInst == "+") && yInst == "=" { // horizontal
		if xInst == "-" {
			start = x2
			end = x1
		} else {
			start = x1
			end = x2
		}
		fmt.Printf("\nHorizontal line")
		for i := start; i <= end; i++ {
			diagram[y1][i]++
		}
	} else if (yInst == "-" || yInst == "+") && xInst == "=" { // vertical
		if yInst == "-" {
			start = y2
			end = y1
		} else {
			start = y1
			end = y2
		}
		fmt.Printf("\nVertical line")
		for i := start; i <= end; i++ {
			diagram[i][x1]++
		}
	} else if (xInst == "+" && yInst == "-") || (xInst == "-" && yInst == "+") { // bot left to top right
		if xInst == "-" {
			start = x2
			end = x1
			aux := y1
			y1 = y2
			y2 = aux
		} else {
			start = x1
			end = x2
		}
		fmt.Printf("\nBot Left to Top Right Diagonal")
		for i := start; i <= end; i++ {
			diagram[y1][i]++
			y1--
		}
	} else if (xInst == "+" && yInst == "+") || (xInst == "-" && yInst == "-") { // top left to bot right
		if xInst == "-" {
			start = x2
			end = x1
			aux := y1
			y1 = y2
			y2 = aux
		} else {
			start = x1
			end = x2
		}
		fmt.Printf("\nTop left to Bot Right Diagonal")
		for i := start; i <= end; i++ {
			diagram[y1][i]++
			y1++
		}
	}
	fmt.Println("")
	return diagram
}

func AddInputsToDiagram(tokens []string, diagram [][]int) [][]int {
	for _, token := range tokens {
		res := regexp.MustCompile(`\d+`)
		n := res.FindAllString(token, -1)

		x1, _ := strconv.Atoi(n[0])
		x2, _ := strconv.Atoi(n[2])
		y1, _ := strconv.Atoi(n[1])
		y2, _ := strconv.Atoi(n[3])

		xInst := GetInstructions(x1, x2)
		yInst := GetInstructions(y1, y2)
		diagram = ApplyInstructions(diagram, xInst, yInst, x1, x2, y1, y2)
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

func TestVents(t *testing.T) {
	tokens := []string{"0,9 -> 5,9", "8,0 -> 0,8", "9,4 -> 3,4", "2,2 -> 2,1", "7,0 -> 7,4",
		"6,4 -> 2,0", "0,9 -> 2,9", "3,4 -> 1,4", "0,0 -> 8,8", "5,5 -> 8,2"}
	diagram := BuildDynamicDiagram(10)
	diagram = AddInputsToDiagram(tokens, diagram)
	overlaps := NumberOfOverlappingPoints(diagram)
	PrintDiagram(diagram)
	fmt.Printf("\nTwo lines overlap on %d points.\n", overlaps)
	if overlaps != 12 {
		t.Errorf("Result is incorrect, expected %d, got %d", 12, overlaps)
	}
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
	diagram = AddInputsToDiagram(tokens, diagram)
	overlaps := NumberOfOverlappingPoints(diagram)
	// PrintDiagram(diagram)
	fmt.Printf("\nTwo lines overlap on %d points.\n", overlaps)
}
