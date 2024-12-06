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
	//logrus.SetLevel(logrus.DebugLevel)
	scanner := bufio.NewScanner(os.Stdin)

	reading_rules := true
	var rules []Rule
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			reading_rules = false
		} else {
			if reading_rules {
				rules = append(rules, createRule(line))
				logrus.Debugf("rule: %+v\n", createRule(line))
			} else {
				pages := parseLine(line)
				update := createUpdate(rules, pages)
				logrus.Debugf("update: %+v\n", update.pages)
				if !update.checkRules() {
					update.sort()
					sum += update.getMiddlePage()
				}
				logrus.Debugf("update: %+v\n", update.pages)
			}
		}
	}

	fmt.Printf("sum of middle pages in correct update %d\n", sum)
}
