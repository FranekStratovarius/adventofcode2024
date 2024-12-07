package main

import (
	"regexp"
	"strconv"
)

type Calibration struct {
	solution int
	numbers  []int
}

var find_number, _ = regexp.Compile(`\d+`)

func parseCalibration(line string) (calibration Calibration) {
	solution_position := find_number.FindStringIndex(line)
	// logrus.Debugf("%d %d | %s\n", solution_position[0], solution_position[1], line[solution_position[1]+2:])
	calibration.solution, _ = strconv.Atoi(line[solution_position[0]:solution_position[1]])
	number_strings := find_number.FindAllString(line[solution_position[1]+2:], -1)
	for _, number_string := range number_strings {
		number, _ := strconv.Atoi(number_string)
		calibration.numbers = append(calibration.numbers, number)
	}
	// logrus.Debugf("%d %v\n", calibration.solution, calibration.operators)
	return
}

func calculateAllSolutions(solution int, number int, numbers []int) bool {
	if len(numbers) == 0 {
		return solution == number
	} else {
		return calculateAllSolutions(
			solution,
			number+numbers[0],
			numbers[1:],
		) ||
			calculateAllSolutions(
				solution,
				number*numbers[0],
				numbers[1:],
			)
	}
}
