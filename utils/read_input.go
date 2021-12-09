package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadInputAsTokens(input string) []string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	
	var tokens []string
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}
	return tokens
}