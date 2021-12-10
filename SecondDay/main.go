package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReturnMultipliedCommands(tokens []string) int{
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
	
	result := ReturnMultipliedCommands(tokens)
	fmt.Println("X and Y coordinates combined: " + strconv.Itoa(result))
}