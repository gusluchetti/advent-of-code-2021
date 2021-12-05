package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var depthArray []int;

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		depthArray = append(depthArray, depth)
	}
	
	var counter int = 0;
	for i:=1; i < len(depthArray); i++ {
		if depthArray[i] > depthArray[i-1] {
			fmt.Printf("Depth %d is bigger than depth %d \n", depthArray[i], depthArray[i-1])
			counter++
			fmt.Printf("Counter increased! It is now at %d \n", counter)
		}
	}

	fmt.Println("Number of times depth increased: " + strconv.Itoa(counter))
}
