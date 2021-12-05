package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
	dat, err := os.ReadFile("input.txt")
	fmt.Print(String(dat))
}
