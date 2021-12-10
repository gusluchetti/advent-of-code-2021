package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	
	res1 := ReturnCommands(tokens)
	res2 := ReturnCommandsWithAim(tokens)
	fmt.Println("Horizontal position and depth coordinates multiplied: " + strconv.Itoa(res1))
	fmt.Println("Same as above, but considering new aim variable: " + strconv.Itoa(res2))
}