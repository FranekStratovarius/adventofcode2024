package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"

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

	var find_number, _ = regexp.Compile(`\d+`)
	var left_column []int
	var right_column []int

	for scanner.Scan() {
		line := scanner.Text()
		// logrus.Debugf("> %s\n", line)

		number_strings := find_number.FindAllString(line, -1)
		left_number, _ := strconv.Atoi(number_strings[1])
		right_number, _ := strconv.Atoi(number_strings[0])
		left_column = append(left_column, left_number)
		right_column = append(right_column, right_number)
	}

	sort.Ints(left_column)
	sort.Ints(right_column)

	sum := 0
	for i, left_number := range left_column {
		sum += int(math.Abs(float64(left_number - right_column[i])))
	}

	fmt.Printf("the total difference is: %d\n", sum)
}
