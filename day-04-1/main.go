package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// logrus.SetLevel(logrus.DebugLevel)
	scanner := bufio.NewScanner(os.Stdin)
	var wordBoard WordBoard
	wordBoard.letters = [][]rune{}
	wordBoard.width = 0
	wordBoard.height = 0

	for scanner.Scan() {
		line := scanner.Text()
		wordBoard.addLine(line)
	}

	sum := 0

	rows := wordBoard.getRows()
	logrus.Debugf("%v\n", rows)
	sum += find_word(rows)

	columns := wordBoard.getColumns()
	logrus.Debugf("%v\n", columns)
	sum += find_word(columns)

	diagonals := wordBoard.getDiagonals()
	logrus.Debugf("%v\n", diagonals)
	sum += find_word(diagonals)

	diagonals_reverse := wordBoard.getDiagonalsReverse()
	logrus.Debugf("%v\n", diagonals_reverse)
	sum += find_word(diagonals_reverse)

	fmt.Printf("found %d matches\n", sum)
}
