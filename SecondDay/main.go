package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReturnCommands(tokens []string) []int{
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

	return commands
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
	
	commands := ReturnCommands(tokens)
	fmt.Println(commands)
}