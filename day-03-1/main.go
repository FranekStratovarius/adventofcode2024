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
	find_numbers, _ := regexp.Compile(`\d+`)
	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Printf("%s\n", line)
		multiplications := find_multiplication.FindAllString(line, -1)
		// fmt.Printf("%v\n", multiplications)
		for _, multiplication := range multiplications {
			numbers := find_numbers.FindAllString(multiplication, -1)
			// fmt.Printf("%+v\n", numbers)
			oins, _ := strconv.Atoi(numbers[0])
			zwoi, _ := strconv.Atoi(numbers[1])
			sum += oins * zwoi
		}
	}

	readFile.Close()
	fmt.Printf("%d\n", sum)
}
