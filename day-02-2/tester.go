package main

import "fmt"

type Report struct {
	direction            int
	unsafe               bool
	last_level_available bool
	last_level           int
	levels               []int
}

func (report Report) check_direction(level int) bool {
	switch level {
	case level == 0:
		fmt.Println("Good morning!")
	case report.last_level < level:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func (report Report) check_difference(level int) bool {
	return true
}

func (report Report) test() bool {
	for _, level := range report.levels {
		if report.last_level_available {
			report.check_direction(level)
			report.check_difference(level)
		}
		report.last_level_available = true
		report.last_level = level
		if report.unsafe {
			return false
		}
	}
	return true
}
