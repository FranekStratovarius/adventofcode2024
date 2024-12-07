package main

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

type Map struct {
	tiles           [][]bool
	tiles_visited   [][]bool
	guard_x         int
	guard_y         int
	guard_direction int
}

var find_obstacles, _ = regexp.Compile(`#`)
var find_guard, _ = regexp.Compile(`\^`)

func (tilemap *Map) addRow(line string) {
	row := make([]bool, len(line))
	for i := 0; i < len(line); i++ {
		row[i] = false
	}
	obstacles := find_obstacles.FindAllStringIndex(line, -1)
	for _, obstacle := range obstacles {
		row[obstacle[0]] = true
	}
	guard := find_guard.FindAllStringIndex(line, -1)
	if len(guard) > 0 {
		tilemap.guard_x = guard[0][0]
		tilemap.guard_y = len(tilemap.tiles)
		tilemap.guard_direction = 0
	}
	tilemap.tiles = append(tilemap.tiles, row)

	row_visited := make([]bool, len(line))
	tilemap.tiles_visited = append(tilemap.tiles_visited, row_visited)
}

func (tilemap Map) print() {
	for y, row := range tilemap.tiles {
		for x, tile := range row {
			if x == tilemap.guard_x && y == tilemap.guard_y {
				switch tilemap.guard_direction {
				case 0:
					// guard looking up
					logrus.Debugf("^")
				case 1:
					// guard looking to the right
					logrus.Debugf(">")
				case 2:
					// guard looking down
					logrus.Debugf("V")
				case 3:
					// guard looking to the left
					logrus.Debugf("<")
				}
			} else {
				if tilemap.tiles_visited[y][x] {
					logrus.Debugf("X")
				} else {
					if tile {
						logrus.Debugf("#")
					} else {
						logrus.Debugf(".")
					}
				}
			}
		}
		logrus.Debugf("\n")
	}
}

func (tilemap Map) countVisited() int {
	visited_tiles := 0
	for _, row := range tilemap.tiles_visited {
		for _, tile := range row {
			if tile {
				visited_tiles++
			}
		}
	}
	return visited_tiles
}
