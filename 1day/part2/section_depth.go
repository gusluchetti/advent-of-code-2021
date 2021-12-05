package main

import (
	"fmt"
)

func GetSlidingWindowsDepthIncreases(depthArray []int) int {
	var counter int = 0;
		for i:=1; i < len(depthArray); i++ {
			if depthArray[i] > depthArray[i-1] {
				fmt.Printf("%d at index %d > %d at index %d\n", depthArray[i], i, depthArray[i-1], i-1)
				counter++
			}

		}
		return counter
}