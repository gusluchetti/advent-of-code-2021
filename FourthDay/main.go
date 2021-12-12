package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func GetDrawnNumbers(tokens []string) []int {
	var drawNumbers []int
	split := strings.Split(tokens[0], ",")
	for _, value := range split {
		number, _ := strconv.Atoi(value)
		drawNumbers = append(drawNumbers, number)
	}

	return drawNumbers
}

func GetBoardsAndStates(tokens []string) ([][][]int, [][][]bool) {
	var boards [][][]int
	var states [][][]bool
	board := make([][]int, 5)
	state := make([][]bool, 5)
	lineNumber := 0

	for i := 2; i < len(tokens); i++ {
		split := strings.Fields(tokens[i])
		if len(split) == 5 {
			board[lineNumber] = make([]int, 5)
			state[lineNumber] = make([]bool, 5)

			for i, value := range split {
				number, _ := strconv.Atoi(value)
				board[lineNumber][i] = number
				state[lineNumber][i] = false
			}
			lineNumber++
		}
		if lineNumber == 5 {
			boards = append(boards, board)
			states = append(states, state)
			lineNumber = 0
			board = make([][]int, 5)
			state = make([][]bool, 5)
		}
	}

	return boards, states
}

func BoardHasWon(state [][]bool) bool {
	victory := false
	count := 0
	// vertical check
	for x := 0; x < len(state); x++ {
		for y := 0; y < len(state); y++ {
			if state[y][x] {
				count++
			}
		}
		if count == 5 {
			victory = true
			break
		} else {
			count = 0
		}
	}
	if victory {
		return true
	}

	victory = false
	count = 0
	// horizontal check
	for y := 0; y < len(state); y++ {
		for x := 0; x < len(state); x++ {
			if state[y][x] {
				count++
			}
		}
		if count == 5 {
			victory = true
			break
		} else {
			count = 0
		}
	}

	return victory
}

func GameLoop(drawnNumbers []int, boards [][][]int, states [][][]bool) (int, int) {
	var lastDrawnNumber int
	var boardIndex int
	var winnersArray []int
	count := 0
	firstScore := 0
	lastScore := 0

numbers:
	for _, current := range drawnNumbers {
		color.Blue("\nDrawn number is %d", current)
		fmt.Printf("Ignoring boards of index %v\n", winnersArray)
	checkBoard:
		for b, board := range boards {
			// we should ignore boards that have already won
			for _, bIndex := range winnersArray {
				if b == bIndex {
					continue checkBoard
				}
			}
			for y, array := range board {
				for x, number := range array {
					if number == current {
						states[b][y][x] = true
						if BoardHasWon(states[b]) {
							color.Magenta("\nBoard %d won!!!", b+1)
							lastDrawnNumber = current
							boardIndex = b
							if len(winnersArray) == 0 {
								firstScore = GetTotalScore(boards[b], states[b], current)
							}
							winnersArray = append(winnersArray, boardIndex)
							if len(winnersArray) == len(boards) {
								break numbers
							}
							count++
						}
					}
				}
			}
			PrintBoardWithHighlightedNumbers(board, states[b])
		}
	}

	fmt.Printf("\nOrder of winner boards: %v", winnersArray)
	lastIndex := winnersArray[len(boards)-1]
	lastScore = GetTotalScore(boards[lastIndex], states[lastIndex], lastDrawnNumber)
	return firstScore, lastScore
}

func GetTotalScore(board [][]int, state [][]bool, lastDrawnNumber int) int {
	score := 0
	for i, array := range board {
		for j, number := range array {
			if !state[i][j] {
				score += number
			}
		}
	}

	fmt.Printf("\nSum of unmarked numbers is %d", score)
	totalScore := score * lastDrawnNumber
	fmt.Printf("\nFinal score is %d * %d = %d\n", score, lastDrawnNumber, totalScore)

	return totalScore
}

func PrintBoardWithHighlightedNumbers(board [][]int, state [][]bool) {
	for i, line := range board {
		fmt.Print(color.WhiteString("["))
		for j, number := range line {
			if state[i][j] {
				color.Set(color.FgYellow)
				fmt.Printf(" %02d ", number)
				color.Unset()
			} else {
				fmt.Printf(" %02d ", number)
			}
		}
		fmt.Print(color.WhiteString("]\n"))
	}
	fmt.Print("\n")
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

	drawnNumbers := GetDrawnNumbers(tokens)
	boards, states := GetBoardsAndStates(tokens)

	firstWinScore, lastWinScore := GameLoop(drawnNumbers, boards, states)

	fmt.Println("\nScore of first board to win is " + color.BlueString(strconv.Itoa(firstWinScore)))
	fmt.Println("\nScore of last board to win is " + color.RedString(strconv.Itoa(lastWinScore)))
}
