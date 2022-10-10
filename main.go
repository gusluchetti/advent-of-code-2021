package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/gusluchetti/advent-of-code/utils"
)

type info struct {
	year int
	day int
	part int
	tail []string
}

func LoadAllInputs() {
	// loading all possible inputs
	years := [7]int{2015, 2016, 2017, 2018, 2019, 2020, 2021}
	for _, y := range years {
		utils.DumpInput(y)
	}
}

func GetSolution(part int) {
	if part == 3 {
		GetSolution(1)
		GetSolution(2)
	} else {
		GetSolution(part)
	}
}

func GetSolutions(info info) {
	years := [7]int{2015, 2016, 2017, 2018, 2019, 2020, 2021}
	for year:=years[0]; year<years[len(years)-1]; year++{
		if info.year != 0 && info.year != year {
			continue
		}
		fmt.Printf("\n-YEAR %d", info.year)
		for day:=1; day<=25; day++ {
			if info.day != 0 && info.day != day {
				continue
			}
			// main_dir/src/YEAR/DAY/YEARDAY.go, run one or both solution
			fmt.Printf("\n>DAY %02d", info.day)
			switch info.part {
			case 3:
				GetSolution(3)
			default:
				GetSolution(info.part)
			}
		}
	}


	fmt.Print("solution")
}

func GetProgress() {
	fmt.Print("progress")
}

func ParseFlags() (info, error) {
	year := 0
	flag.IntVar(&year, "y", 0, "year as YYYY")
	day := 0
	flag.IntVar(&day, "d", 0, "day as (zero padded) DD")
	part := 3
	flag.IntVar(&part, "p", 3, "single (1 or 2) or both (3) parts")
	flag.Parse()

	if year >= 2022 {
		return info{}, errors.New("Invalid year!")
	}
	if day < 1 || day > 25 {
		return info{}, errors.New("Invalid day! (Should be between 1 and 25)")
	}
	if part < 1 || part > 3 {
		return info{}, errors.New("Invalid part! (Can be either 1, 2 or 3)")
	}

	fmt.Println("Printing all flags")
	fmt.Println(year, day, part, flag.Args())
	tail := flag.Args()

	return info{
		year: year,
		day: day,
		part: part,
		tail: tail,
	}, nil
}

func main() {
	// LoadAllInputs()
	// fmt.Println("All inputs are loaded.")

	info, err := ParseFlags()
	if err != nil {

	}

	if info.tail[0] == "solve" {
		GetSolutions(info)
	} else if info.tail[0] == "progress"{
		GetProgress()
	}
}
