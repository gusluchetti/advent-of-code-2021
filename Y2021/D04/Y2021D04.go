package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

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

func TestBingo (t * testing.T) {
	tokens := []string{
		"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1", "",
		"22 13 17 11  0", "8  2 23  4 24", "21  9 14 16  7", "6 10  3 18  5", "1 12 20 15 19", "",
		"3 15  0  2 22", "9 18 13 17  5", "19  8  7 25 23", "20 11 10 24  4", "14 21 16 12  6", "",
		"14 21 17 24  4", "10 16 15  9 19", "18  8 23 26 20", "22 11 13  6  5", "2  0 12  3  7" }
	boards, states := GetBoardsAndStates(tokens)
	firstScore, lastScore := GameLoop(GetDrawnNumbers(tokens), boards, states)

	if firstScore != 4512 {
		t.Errorf("First score was incorrect, go: %d, want: %d", firstScore, 4512)
	} else if lastScore != 1924 {
		t.Errorf("Last score was incorrect, go: %d, want: %d", lastScore, 1924)
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

	drawnNumbers := GetDrawnNumbers(tokens)
	boards, states := GetBoardsAndStates(tokens)

	firstWinScore, lastWinScore := GameLoop(drawnNumbers, boards, states)

	fmt.Println("\nScore of first board to win is " + color.BlueString(strconv.Itoa(firstWinScore)))
	fmt.Println("\nScore of last board to win is " + color.RedString(strconv.Itoa(lastWinScore)))
}
