package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	find_multiplication, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	find_enable, _ := regexp.Compile(`do\(\)`)
	find_disable, _ := regexp.Compile(`don't\(\)`)
	find_numbers, _ := regexp.Compile(`\d+`)
	enabled := true
	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		finished := false
		for !finished {
			//fmt.Print("loop\n")
			if enabled {
				disable_position := find_disable.FindStringIndex(line)
				multiplication_position := find_multiplication.FindStringIndex(line)

				if multiplication_position != nil && (disable_position == nil || multiplication_position[0] < disable_position[0]) {
					numbers := find_numbers.FindAllString(line[multiplication_position[0]:multiplication_position[1]], -1)
					//fmt.Printf("%s\n", numbers)
					oins, _ := strconv.Atoi(numbers[0])
					zwoi, _ := strconv.Atoi(numbers[1])
					sum += oins * zwoi
					line = line[multiplication_position[1]:]
				} else if disable_position != nil {
					line = line[disable_position[1]:]
					//fmt.Printf("disable: %s\n", line)
					enabled = false
				} else {
					finished = true
				}
			} else {
				//fmt.Print("not enabled\n")
				enable_position := find_enable.FindStringIndex(line)
				if enable_position != nil {
					//fmt.Printf("%d, %d\n", enable_position[0], enable_position[1])
					line = line[enable_position[1]:]
					//fmt.Printf("enable: %s\n", line)
					enabled = true
				} else {
					finished = true
				}
			}
		}
	}

	readFile.Close()
	fmt.Printf("%d\n", sum)
}
