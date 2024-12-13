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

	stones := makeStones(input)
	print(stones)
	for i := 0; i < 75; i++ {
		logrus.Debugf("blink: %d, ", i+1)
		stones = blink(stones)
		logrus.Debugf("number of stones: %d\n", count(stones))
		print(stones)
	}
	logrus.Debugf("\n")

	logrus.Debugf("number of stones: %d\n", count(stones))
}
