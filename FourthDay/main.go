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
		if (lineNumber == 5) {
			boards = append(boards, board)
			states = append(states, state)
			lineNumber = 0
			board = make([][]int, 5)
			state = make([][]bool, 5)
		}
	}

	return boards, states
}

func CheckBoardState(state [][]bool) bool {
	victoryAchieved := true
	for i, array := range state {
		for j, single := range array {
			length := len(array) - 1
			half := length / 2
			if (single) {
				// check if win might be diagonal
				if (i == j && i + j == length) {
					if ((i == half) || (i < half && j < half) || (i > half && j > half)) { // top left to bot right diag
						for b := 0; b<=length; b++ {
							y := 0
							if (!state[y][b]) {
								victoryAchieved = false
								break
							}
							y++
						}
						if (victoryAchieved) {
							return true
						}
					} else if (i == half){ // bot left to top right diag
						for b := 0; b<=length; b++ {
							y := length
							if (!state[y][b]) {
								victoryAchieved = false
								break
							}
							y--
						}
						if (victoryAchieved) {
							return true
						}
					}
				}
				// check for horizontal win
				for _, n := range state[i] {
					if (!n) {
						victoryAchieved = false
						break
					}
				}
				if (victoryAchieved) {
					return true
				}
				// check for vertical win
				for _, m := range state {
					if(!m[j]) {
						victoryAchieved = false
						break
					}
				}
				if (victoryAchieved) {
					return true
				}
			}
		}
	}
	return victoryAchieved
}

func GameLoop(drawnNumbers []int, boards [][][]int, states[][][]bool) int {
	var lastDrawnNumber int
	var boardIndex int
	gameFinished := false
	score := 0

	for i, current := range drawnNumbers {
		color.Cyan("\n#%d Number is %d", i+1, current)
		for b, board := range boards {
			for y, array := range board {
				for x, number := range array {
					if number == current {
						states[b][y][x] = true
						PrintBoardWithHighlightedNumbers(board, states[b])
						// checking if this board won the game
						if (CheckBoardState(states[b])) {
							color.Magenta("\nBoard %d won!!!", b+1)
							lastDrawnNumber = current
							boardIndex = b
							gameFinished = true
							break
						}
					}				
				}
			}
		}

		if (gameFinished) {
			score = GetWinningScore(boards[boardIndex], states[boardIndex], lastDrawnNumber)
			break
		}
	}

	return score
}

func GetWinningScore(board [][]int, state [][]bool, lastDrawnNumber int) int {
	score := 0
	for i, array := range board {
		for j, number := range array {
			if (!state[i][j]) {
				score += number
			}
		}
	}
	return score * lastDrawnNumber
}

func PrintBoardWithHighlightedNumbers(board [][]int, state [][]bool) {
	for i, line := range board {
		fmt.Print(color.WhiteString("["))
		for j, number := range line {
			if(state[i][j]) {
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

	boards, states := GetBoardsAndStates(tokens)
	score := GameLoop(GetDrawnNumbers(tokens), boards, states)
	fmt.Printf("\nScore for winning board was %d", score)
}