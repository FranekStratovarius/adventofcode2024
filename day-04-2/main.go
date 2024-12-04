package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type PlainFormatter struct{}

func (formatter *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf(entry.Message)), nil
}

func main() {
	logrus.SetFormatter(&PlainFormatter{})
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

	sum := wordBoard.checkMatches()
	wordBoard.print()

	fmt.Printf("found %d matches\n", sum)
}
