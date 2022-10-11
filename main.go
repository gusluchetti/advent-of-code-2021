package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/gusluchetti/advent-of-code/utils"
)

type flags struct {
	year int
	day int
	tail []string
}

func parseFlags() (flags, error) {
	year := 0
	flag.IntVar(&year, "y", 0, "year as YYYY")
	day := 0
	flag.IntVar(&day, "d", 0, "day as (zero padded) DD")
	flag.Parse()
	tail := flag.Args()

	if  year > 2021 || (year < 2015 && year != 0) {
		return flags{}, errors.New("Invalid year!")
	}
	if  day > 25 || (day < 1 && day != 0) {
		return flags{}, errors.New("Invalid day! (Should be between 1 and 25)")
	}

	return flags{
		year: year,
		day: day,
		tail: tail,
	}, nil
}

func getDaySolutions(path string) {
	cmd := exec.Command("go", "run", path)
	output, err := cmd.CombinedOutput()
	utils.Check(err)
	fmt.Printf("\n%s", string(output))
}

func getAllSolutions(flags flags) {
	fmt.Print("Starting to get solutions...")
	years := [7]int{2015, 2016, 2017, 2018, 2019, 2020, 2021}
	for year:=years[0]; year<=years[len(years)-1]; year++{
		if flags.year != 0 && flags.year != year {
			continue
		}
		fmt.Printf("\n- YEAR %d -", year)
		for day:=1; day<=25; day++ {
			if flags.day != 0 && flags.day != day {
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
				getDaySolutions(targetFile)
			}
		}
	}
}

func getProgress(flags flags) {
	// TODO: implement progress status
	fmt.Print("progress")
}

func main() {
	// loading all possible inputs
	years := [7]int{2015, 2016, 2017, 2018, 2019, 2020, 2021}
	for _, y := range years {
		utils.DumpInput(y)
	}

	flags, err := parseFlags()
	utils.Check(err)

	mode := flags.tail[0]
	switch mode {
	case "solve":
		getAllSolutions(flags)
	case "progress":
		getProgress(flags)
	}
}
