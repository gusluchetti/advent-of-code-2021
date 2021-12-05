package main

func GetSections(depthArray []int) []int {
	sections := len(depthArray) % 3
	var sectionsArray []int
		for i:=1; i < sections; i++ {
			sectionsArray = append(sectionsArray, (depthArray[i] + depthArray[i+1] + depthArray[i+2]))
		}
		return sectionsArray
}