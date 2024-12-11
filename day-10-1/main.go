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
	logrus.SetLevel(logrus.DebugLevel)
	scanner := bufio.NewScanner(os.Stdin)

	var hikingMap HikingMap
	for scanner.Scan() {
		line := scanner.Text()
		// logrus.Debugf("> %s\n", line)
		hikingMap.addRow(line)
	}

	// hikingMap.print()

	trailheadRatingSum := hikingMap.findHikingTrails()

	fmt.Printf("summed scores of trailheads: %d\n", trailheadRatingSum)
}
