package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, err := os.Open("test_input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	defer inputFile.Close()
	fileScanner.Split(bufio.ScanLines)

	var reports []Report

	r, _ := regexp.Compile(`[0123456789]+`)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Printf("%s\n", line)
		levels := r.FindAllString(line, -1)
		fmt.Printf("%v\n", levels)
		var report []int
		for _, level := range levels {
			int_level, _ := strconv.Atoi(level)
			report = append(report, int_level)
		}
		reports = append(reports, Report{
			direction:            0,
			unsafe:               false,
			last_level_available: false,
			last_level:           0,
			levels:               report,
		})
	}

	safe := 0
	for _, report := range reports {
		if report.test() {
			safe++
		}
	}

	fmt.Printf("%d\n", safe)
}
