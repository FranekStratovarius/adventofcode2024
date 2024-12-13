package main

import (
	"math"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

func print(stones map[int]int) {
	for number, stoneCount := range stones {
		logrus.Debugf("%d: %d, ", number, stoneCount)
	}
	logrus.Debugf("\n")
}

var find_number, _ = regexp.Compile(`\d+`)

func makeStones(input string) (stones map[int]int) {
	stones = make(map[int]int)
	numberStrings := find_number.FindAllString(input, -1)
	for _, numberString := range numberStrings {
		number, _ := strconv.Atoi(numberString)
		stones[number]++
	}
	return
}

func blink(stones map[int]int) (nextStones map[int]int) {
	nextStones = make(map[int]int)
	for number, stoneCount := range stones {
		if number == 0 {
			nextStones[1] += stoneCount
		} else {
			digitsCount := getDigits(number)
			if digitsCount%2 == 0 {
				left, right := splitInteger(number, digitsCount)
				nextStones[left] += stoneCount
				nextStones[right] += stoneCount
			} else {
				nextStones[number*2024] += stoneCount
			}
		}
	}
	return
}

func getDigits(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}

func splitInteger(i int, digitsCount int) (left int, right int) {
	leftScalar := int(math.Pow(10.0, float64(digitsCount/2)))
	left = i / leftScalar
	right = i - left*leftScalar
	return
}

func count(stones map[int]int) (count int) {
	for _, stoneCount := range stones {
		count += stoneCount
	}
	return
}
