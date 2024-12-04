package main

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

type WordBoard struct {
	letters       [][]rune
	width, height int
}

func (wordBoard *WordBoard) addLine(line string) {
	var row []rune
	for _, v := range line {
		row = append(row, v)
	}
	wordBoard.letters = append(wordBoard.letters, row)
	if len(line) > wordBoard.width {
		wordBoard.width = len(line)
	}
	wordBoard.height++
}

func (wordBoard WordBoard) getRows() []string {
	var rows []string
	for y := 0; y < wordBoard.height; y++ {
		row := ""
		for x := 0; x < wordBoard.width; x++ {
			row += string(wordBoard.letters[y][x])
		}
		rows = append(rows, row)
	}
	return rows
}

func (wordBoard WordBoard) getColumns() []string {
	var rows []string
	for x := 0; x < wordBoard.width; x++ {
		row := ""
		for y := 0; y < wordBoard.height; y++ {
			row += string(wordBoard.letters[y][x])
		}
		rows = append(rows, row)
	}
	return rows
}

func (wordBoard WordBoard) getDiagonals() []string {
	var rows []string

	for x := 0; x < wordBoard.width; x++ {
		row := ""
		for y := 0; y < wordBoard.height && y+x < wordBoard.width; y++ {
			row += string(wordBoard.letters[y][x+y])
		}
		rows = append(rows, row)
	}

	for y := 1; y < wordBoard.height; y++ {
		row := ""
		for x := 0; x < wordBoard.width && x+y < wordBoard.height; x++ {
			row += string(wordBoard.letters[x+y][x])
		}
		rows = append(rows, row)
	}

	return rows
}

func (wordBoard WordBoard) getDiagonalsReverse() []string {
	var rows []string

	for x := 0; x < wordBoard.width; x++ {
		row := ""
		for y := 0; y < wordBoard.height && y+x < wordBoard.width; y++ {
			row += string(wordBoard.letters[wordBoard.height-y-1][x+y])
		}
		rows = append(rows, row)
	}

	for y := 1; y < wordBoard.height; y++ {
		row := ""
		for x := 0; x < wordBoard.width && x+y < wordBoard.height; x++ {
			row += string(wordBoard.letters[wordBoard.height-y-1-x][x])
		}
		rows = append(rows, row)
	}

	return rows
}

var find_xmas, _ = regexp.Compile(`XMAS`)
var find_samx, _ = regexp.Compile(`SAMX`)

func find_word(rows []string) int {
	sum := 0
	for _, row := range rows {
		matches := find_xmas.FindAllStringIndex(row, -1)
		logrus.Debugf("\t%v\n", matches)
		sum += len(matches)
		matches = find_samx.FindAllStringIndex(row, -1)
		logrus.Debugf("\t%v\n", matches)
		sum += len(matches)
	}
	return sum
}
