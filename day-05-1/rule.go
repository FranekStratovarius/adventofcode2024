package main

import (
	"regexp"
	"strconv"
)

type Rule struct {
	first_page  int
	second_page int
}

var find_rule, _ = regexp.Compile(`\d+`)

func createRule(line string) Rule {
	pages := find_rule.FindAllString(line, -1)
	first_page, _ := strconv.Atoi(pages[0])
	second_page, _ := strconv.Atoi(pages[1])

	return Rule{
		first_page:  first_page,
		second_page: second_page,
	}
}

func (rule Rule) checkRule(pages []int) bool {
	page_hit := false
	for _, page := range pages {
		if page == rule.second_page {
			page_hit = true
		}
		if page_hit && page == rule.first_page {
			return false
		}
	}
	return true
}
