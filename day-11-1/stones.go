package main

import (
	"math"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Stone struct {
	value int
	next  *Stone
}

var find_number, _ = regexp.Compile(`\d+`)

func (stone *Stone) print() {
	lastStone := stone
	for lastStone != nil {
		logrus.Debugf("%d, ", lastStone.value)
		lastStone = lastStone.next
	}
	logrus.Debugf("\n")
}

func makeStones(input string) (firstStone Stone) {
	var lastStone *Stone
	numberStrings := find_number.FindAllString(input, -1)
	for _, numberString := range numberStrings {
		number, _ := strconv.Atoi(numberString)
		if lastStone == nil {
			lastStone = &firstStone
		} else {
			lastStone.next = &Stone{}
			lastStone = lastStone.next
		}
		lastStone.value = number

	}
	return
}

func (stone *Stone) blink() {
	lastStone := stone
	for lastStone != nil {
		if lastStone.value == 0 {
			lastStone.value = 1
			lastStone = lastStone.next
		} else {
			digitsCount := getDigits(lastStone.value)
			if digitsCount%2 == 0 {
				// replace by 2 stones
				left, right := splitInteger(lastStone.value, digitsCount)
				lastStone.value = left
				nextStone := lastStone.next
				lastStone.next = &Stone{
					value: right,
					next:  nextStone,
				}
				lastStone = nextStone
			} else {
				lastStone.value *= 2024
				lastStone = lastStone.next
			}
		}
	}
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

func (stone *Stone) count() (count int) {
	lastStone := stone
	for lastStone != nil {
		count++
		lastStone = lastStone.next
	}
	return
}
