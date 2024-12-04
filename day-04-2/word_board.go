package main

import "github.com/sirupsen/logrus"

type WordBoard struct {
	letters       [][]rune
	hit           [][]bool
	width, height int
}

func (wordBoard *WordBoard) addLine(line string) {
	var row []rune
	for _, v := range line {
		row = append(row, v)
	}
	wordBoard.letters = append(wordBoard.letters, row)
	wordBoard.hit = append(wordBoard.hit, make([]bool, len(line)))
	if len(line) > wordBoard.width {
		wordBoard.width = len(line)
	}
	wordBoard.height++
}

func (wordBoard WordBoard) print() {
	for y := 0; y < wordBoard.height; y++ {
		for x := 0; x < wordBoard.width; x++ {
			if wordBoard.hit[y][x] {
				logrus.Debugf("%s", string(wordBoard.letters[y][x]))
			} else {
				logrus.Debugf(".")
			}
		}
		logrus.Debug("\n")
	}
}

func (wordBoard *WordBoard) checkMatch(x, y int) bool {
	if
	// check center
	(wordBoard.letters[y][x] == 'A') && (
	// check top left to bottom right
	(
	// if top left is a M
	((wordBoard.letters[y-1][x-1] == 'M') &&
		// then bottom right has to be a S
		(wordBoard.letters[y+1][x+1] == 'S')) ||
		// or
		// if top left is a S
		((wordBoard.letters[y-1][x-1] == 'S') &&
			// then bottom right has to be a M
			(wordBoard.letters[y+1][x+1] == 'M'))) && (
	// if top right is a M
	((wordBoard.letters[y-1][x+1] == 'M') &&
		// then bottom left has to be a S
		(wordBoard.letters[y+1][x-1] == 'S')) ||
		// or
		// if top right is a S
		((wordBoard.letters[y-1][x+1] == 'S') &&
			// then bottom left has to be a M
			(wordBoard.letters[y+1][x-1] == 'M')))) {
		wordBoard.hit[y+1][x-1] = true
		wordBoard.hit[y+1][x+1] = true
		wordBoard.hit[y][x] = true
		wordBoard.hit[y-1][x-1] = true
		wordBoard.hit[y-1][x+1] = true
		return true
	}
	return false
}

func (wordBoard *WordBoard) checkMatches() int {
	sum := 0
	for y := 1; y < wordBoard.height-1; y++ {
		for x := 1; x < wordBoard.width-1; x++ {
			if wordBoard.checkMatch(x, y) {
				sum++
			}
		}
	}
	return sum
}
