package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetPowerConsumption(tokens []string) int64 {
	bitSize := len(strings.Split(tokens[1], ""))
	
	gammaBits := make([]string, bitSize)
	epsilonBits := make([]string, bitSize)
	temp := make([]int, bitSize)
	for _, token := range tokens {
		split := strings.Split(token, "")
		for i, char := range split {
			bit, err := strconv.Atoi(char)
			if err != nil { log.Fatal(err) }
			temp[i] += bit
		}
	}

	for i, bit := range temp {
		size := len(tokens)
		// most common bit is 1
		if bit > size / 2 {
			gammaBits[i] = "1"
			epsilonBits[i] = "0"
		} else { // most common bit is 0
			gammaBits[i] = "0"
			epsilonBits[i] = "1"
		}
	}
	gamma, _ := strconv.ParseInt(strings.Join(gammaBits, ""), 2, 64)
	epsilon, _ := strconv.ParseInt(strings.Join(epsilonBits, ""), 2, 64) 
	fmt.Println(gamma, epsilon)
	return gamma * epsilon
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
	
	res1 := GetPowerConsumption(tokens)
	fmt.Println("Power Consumption: " + strconv.Itoa(int(res1)))
}