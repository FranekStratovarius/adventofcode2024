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

	var calibrations []Calibration

	for scanner.Scan() {
		line := scanner.Text()
		// logrus.Debugf("> %s\n", line)
		calibrations = append(calibrations, parseCalibration(line))
	}

	result := 0
	for _, calibration := range calibrations {
		if calculateAllSolutions(calibration.solution, calibration.numbers[0], calibration.numbers[1:]) {
			// if calibration.checkPossible() {
			result += calibration.solution
		}
	}

	fmt.Printf("total calibration result: %d\n", result)
}
