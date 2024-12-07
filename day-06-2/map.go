package main

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

type Map struct {
	tiles            [][]bool
	tiles_visited    [][]bool
	tiles_obstructed [][]bool
	guard_x          int
	guard_y          int
	guard_direction  int
	guard_memory     [][3]int
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
	row_obstructed := make([]bool, len(line))
	tilemap.tiles_visited = append(tilemap.tiles_visited, row_visited)
	tilemap.tiles_obstructed = append(tilemap.tiles_obstructed, row_obstructed)
}

func (tilemap Map) print() {
	for _, visited_position := range tilemap.guard_memory {
		tilemap.tiles_visited[visited_position[1]][visited_position[0]] = true
		logrus.Debugf("%v\n", visited_position)
	}
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
				if tilemap.tiles_obstructed[y][x] {
					logrus.Debugf("O")
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
		}
		logrus.Debugf("\n")
	}
}

func (tilemap Map) copy() (new_tilemap Map) {
	new_tilemap.tiles = copyDoubleSlice(tilemap.tiles)
	new_tilemap.tiles_visited = copyDoubleSlice(tilemap.tiles_visited)
	new_tilemap.tiles_obstructed = copyDoubleSlice(tilemap.tiles_obstructed)
	new_tilemap.guard_x = tilemap.guard_x
	new_tilemap.guard_y = tilemap.guard_y
	new_tilemap.guard_direction = tilemap.guard_direction
	return
}

func copyDoubleSlice(tilemap [][]bool) (new_tiles [][]bool) {
	for _, row := range tilemap {
		var new_row []bool
		for _, tile := range row {
			new_row = append(new_row, tile)
		}
		new_tiles = append(new_tiles, new_row)
	}
	return
}

func (tilemap Map) countObstructed() int {
	obstructed_tiles := 0
	for _, row := range tilemap.tiles_obstructed {
		for _, tile := range row {
			if tile {
				obstructed_tiles++
			}
		}
	}
	return obstructed_tiles
}

func (tilemap Map) getNextPosition() (pos_x int, pos_y int, left bool) {
	pos_x = tilemap.guard_x
	pos_y = tilemap.guard_y
	left = false
	switch tilemap.guard_direction {
	case 0:
		pos_y--
	case 1:
		pos_x++
	case 2:
		pos_y++
	case 3:
		pos_x--
	}

	// if guard left the field, return false
	if pos_x < 0 || pos_y < 0 ||
		pos_x > len(tilemap.tiles[0])-1 ||
		pos_y > len(tilemap.tiles)-1 {
		left = true
	}
	return
}
