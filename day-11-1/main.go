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

	var input string
	for scanner.Scan() {
		line := scanner.Text()
		logrus.Debugf("> %s\n", line)
		input = line
	}

	// input := "125 17"

	stone := makeStones(input)
	stone.print()
	for i := 0; i < 25; i++ {
		stone.blink()
		// stone.print()
	}
	logrus.Debugf("number of stones: %d\n", stone.count())
}
