package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	utils "github.com/gusluchetti/advent-of-code/utils"
)

type info struct {
	year int
	day int
	tail []string
}

func LoadAllInputs() {
	// loading all possible inputs
	years := [7]int{2015, 2016, 2017, 2018, 2019, 2020, 2021}
	for _, y := range years {
		utils.DumpInput(y)
	}
}

func GetDaySolutions(path string) {
	cmd := exec.Command("go", "run", path)
	output, err := cmd.CombinedOutput()
	utils.Check(err)
	fmt.Printf("\n%s", string(output))
}

func GetAllSolutions(info info) {
	fmt.Print("Starting to get solutions...")
	years := [7]int{2015, 2016, 2017, 2018, 2019, 2020, 2021}
	for year:=years[0]; year<=years[len(years)-1]; year++{
		if info.year != 0 && info.year != year {
			continue
		}
		fmt.Printf("\n- YEAR %d -", year)
		for day:=1; day<=25; day++ {
			if info.day != 0 && info.day != day {
				continue
			}

			path, err := os.Getwd()
			utils.Check(err)
			targetDir := fmt.Sprintf("%s/src/Y%d/D%02d", path, year, day)
			targetFile := utils.GetTargetFile(targetDir, year, day) + ".go"

			if _, err := os.Stat(targetFile); errors.Is(err, os.ErrNotExist) {
				// fmt.Printf("\nNo solutions for day %d", day)
			} else {
				fmt.Printf("\n# DAY %02d", day)
				GetDaySolutions(targetFile)
			}
		}
	}
}

func GetProgress() {
	// TODO: implement progress status
	fmt.Print("progress")
}

func ParseFlags() (info, error) {
	year := 0
	flag.IntVar(&year, "y", 0, "year as YYYY")
	day := 0
	flag.IntVar(&day, "d", 0, "day as (zero padded) DD")
	flag.Parse()

	if  year > 2021 || (year < 2015 && year != 0) {
		return info{}, errors.New("Invalid year!")
	}
	if  day > 25 || (day < 1 && day != 0) {
		return info{}, errors.New("Invalid day! (Should be between 1 and 25)")
	}

	fmt.Println("Printing all flags")
	fmt.Println(year, day, flag.Args())
	tail := flag.Args()

	return info{
		year: year,
		day: day,
		tail: tail,
	}, nil
}

func main() {
	// LoadAllInputs()
	// fmt.Println("All inputs are loaded.")

	info, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	if info.tail[0] == "solve" {
		GetAllSolutions(info)
	} else if info.tail[0] == "progress"{
		GetProgress()
	}
}
