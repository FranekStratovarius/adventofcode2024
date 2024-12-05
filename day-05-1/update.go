package main

import (
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Update struct {
	rules []Rule
	pages []int
}

var find_page, _ = regexp.Compile(`\d+`)

func parseLine(line string) []int {
	var pages []int
	logrus.Debugf("\t\"%s\"\n", line)
	pages_strings := find_page.FindAllString(line, -1)
	logrus.Debugf("\t%v\n", pages_strings)
	for _, page_string := range pages_strings {
		page, _ := strconv.Atoi(page_string)
		pages = append(pages, page)
	}
	return pages
}

func createUpdate(rules []Rule, pages []int) Update {
	return Update{
		rules: rules,
		pages: pages,
	}
}

func (update Update) checkRules() bool {
	for _, rule := range update.rules {
		if !rule.checkRule(update.pages) {
			return false
		}
	}
	return true
}

func (update Update) getMiddlePage() int {
	return update.pages[len(update.pages)/2]
}
