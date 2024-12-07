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

	tilemap := Map{}

	for scanner.Scan() {
		line := scanner.Text()
		logrus.Printf("> %s\n", line)
		tilemap.addRow(line)
	}

	// tilemap.print()
	// logrus.Debugf("\n--------------\n\n")
	for tilemap.simulateStep() {
		// tilemap.print()
		// logrus.Debugf("\n--------------\n\n")
		// time.Sleep(20 * time.Millisecond)
	}

	visited_tiles := tilemap.countVisited()

	fmt.Printf("visited positions %d\n", visited_tiles)
}
