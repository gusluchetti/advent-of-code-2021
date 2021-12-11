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
	boardNumber := 0
	
	for i := 2; i < len(tokens); i++ {
		split := strings.Fields(tokens[i])
		fmt.Printf("\nBoard Number: %d, Line Number: %d", boardNumber, lineNumber)
		fmt.Printf("\nCurrent split: %s [Length %d]\n", split, len(split))
		
		if len(split) == 5 {
			board[lineNumber] = make([]int, 5)
			state[lineNumber] = make([]bool, 5)

			for i, value := range split {
				number, _ := strconv.Atoi(value)
				board[lineNumber][i] = number
				state[lineNumber][i] = false
			}
			lineNumber++
		} else {
			fmt.Printf("\nFinished board and state: %v, %v \n", board, state)
			boards = append(boards, board)
			states = append(states, state)

			boardNumber++
			lineNumber = 0
		}
	}

	return boards, states
}

func CheckWinConditions() bool {
	isGameFinished := false


	return isGameFinished
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

	color.Blue("Game is about to start!")
	// game loop
	isGameFinished := false
	for isGameFinished != true {

		isGameFinished = CheckWinConditions()
	}
}