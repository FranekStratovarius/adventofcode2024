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

	var cityMap Map

	for scanner.Scan() {
		line := scanner.Text()
		// logrus.Debugf("> %s\n", line)
		var row []rune
		for _, tile := range line {
			row = append(row, tile)
		}
		cityMap.addRow(row)
	}

	cityMap.createAntinodeMap()

	cityMap.print()

	antinodesCount := cityMap.countAntinodes()
	fmt.Printf("%d antinodes found\n", antinodesCount)
}
