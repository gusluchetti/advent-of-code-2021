package main

func GetSections(depthArray []int) []int {
	maxLength := len(depthArray) - 2
	var sectionsArray []int
		for i:=0; i < maxLength; i++ {
			sectionsArray = append(sectionsArray, (depthArray[i] + depthArray[i+1] + depthArray[i+2]))
		}

		return sectionsArray
}