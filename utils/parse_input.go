package utils

import (
	"os"
	"log"
	"bufio"
)

func ParseInput(filepath string) []string {
	file, err := os.Open(filepath)
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
