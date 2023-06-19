package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func GetMostCommonBitsPerColumn(bitCount []int, tokens []string) []string {
	for _, token := range tokens {
		split := strings.Split(token, "")
		for i, char := range split {
			bit, err := strconv.Atoi(char)
			if err != nil { log.Fatal(err) }
			bitCount[i] += bit
		}
	}

	var commonBits []string
	for _, value := range bitCount {
		if (value > len(tokens) / 2) {
			commonBits = append(commonBits, "1")
		} else {
			commonBits = append(commonBits, "0")
		}
	}
    return commonBits
}

func BuildNumberFromBitArray(bitArray []string, bitSize int, bitsQty int, addMostCommonBit bool) []string {
	listBits := make([]string, bitSize)
	for i, bit := range bitArray {
		switch bit {
		case "1":
			if addMostCommonBit {
				listBits[i] = "1"
			} else {
				listBits[i] = "0"
			}
		case "0":
			if addMostCommonBit {
				listBits[i] = "0"
			} else {
				listBits[i] = "1"
			} 
		}
	}
	return listBits
}

func FilterTokens(tokens []string, NeedMostCommon bool) int64 {
	var validTokens []string
	
	for i := 0; i < len(strings.Split(tokens[0], "")); i++ {
		validTokens = nil
		bits := make([]int, 2)
		for _, token := range tokens {
			split := strings.Split(token, "")
			temp, _ := strconv.Atoi(split[i])
			if temp == 0 {
				bits[0] += 1
			} else {
				bits[1] += 1
			}
		}

		var bit bool
		if bits[0] == bits[1] {
			bit = true
		} else if bits[1] > bits[0] {
			bit = true
		} else {
			bit = false
		}
		if !NeedMostCommon { bit = !bit }
		
		var commonBit string
		if (bit) { 
			commonBit = "1"
		} else { commonBit = "0" }

		for _, token := range tokens {
			split := strings.Split(token, "")
			if split[i] == commonBit {
				validTokens = append(validTokens, token)
			}
		}

		fmt.Println(i, bits, commonBit, tokens, validTokens)
		if len(validTokens) == 1 {
			break
		}

		tokens = validTokens
	}

	res, _ := strconv.ParseInt(validTokens[0], 2, 64)
	fmt.Println(res)
	return res
}

func GetGasRatings(tokens []string) (int64, int64) {
	o2 := FilterTokens(tokens, true)
	co2 := FilterTokens(tokens, false)
	return o2, co2
}

func GetPowerConsumption(tokens []string, bitArray []string, bitSize int) int64 {
	gammaBits := BuildNumberFromBitArray(bitArray, bitSize, len(tokens), true)
	epsilonBits := BuildNumberFromBitArray(bitArray, bitSize, len(tokens), false)
	gamma, _ := strconv.ParseInt(strings.Join(gammaBits, ""), 2, 64)
	epsilon, _ := strconv.ParseInt(strings.Join(epsilonBits, ""), 2, 64) 
	return gamma * epsilon
}

func GetLifeRating(tokens []string) int64 {
 	o2, co2 := GetGasRatings(tokens)
	return o2 * co2
}

func TestPowerConsumption (t * testing.T) {
	tokens := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	bitSize := len(strings.Split(tokens[1], ""))
	mostCommonBitsPerColumn := GetMostCommonBitsPerColumn(make([]int, bitSize), tokens)
	result := GetPowerConsumption(tokens, mostCommonBitsPerColumn, bitSize)

	if result != 198 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", result, 198)
	}
}

func TestLifeSupportRating (t * testing.T) {
	tokens := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	result := GetLifeRating(tokens)

	if result != 230 {
		t.Errorf("Counter was incorrect, go: %d, want: %d", result, 230)
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

	bitSize := len(strings.Split(tokens[1], ""))
	mostCommonBitsPerColumn := GetMostCommonBitsPerColumn(make([]int, bitSize), tokens)

	powerConsumption := GetPowerConsumption(tokens, mostCommonBitsPerColumn, bitSize)
	lifeRating := GetLifeRating(tokens)	

	fmt.Println("Power Consumption: " + strconv.Itoa(int(powerConsumption)))
	fmt.Println("Life Support Rating: " + strconv.Itoa(int(lifeRating)))
}
