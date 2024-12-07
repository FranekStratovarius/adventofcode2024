package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func (tilemap *Map) simulateStep() bool {
	// get next tile
	guard_x, guard_y, left := tilemap.getNextPosition()

	// if guard left the field, return false
	if left {
		return false
	}

	// if len(tilemap.guard_memory) == 0 || !(tilemap.guard_x == tilemap.guard_memory[len(tilemap.guard_memory)-1][0] && tilemap.guard_y == tilemap.guard_memory[len(tilemap.guard_memory)-1][1]) {
	// }
	tilemap.guard_memory = append(tilemap.guard_memory, [3]int{tilemap.guard_x, tilemap.guard_y, tilemap.guard_direction})

	// if guard hits a obstacle, rotate
	if tilemap.tiles[guard_y][guard_x] {
		tilemap.guard_direction = (tilemap.guard_direction + 1) % 4
		// otherwise just walk
	} else {
		tilemap.guard_x = guard_x
		tilemap.guard_y = guard_y
	}
	return true
}

func (tilemap Map) checkObstacle() {
	new_tilemap := tilemap.copy()
	obstacle_x, obstacle_y, left := tilemap.getNextPosition()

	// if the obstacle would be out of the map, just return
	if left {
		return
	}
	// if there is already an obstacle, dont do anything
	// if tilemap.tiles[obstacle_y][obstacle_x] {
	// 	return
	// }
	// dont check for obstacle if it would have blocked the already used path
	for _, visited_position := range tilemap.guard_memory {
		if visited_position[0] == obstacle_x &&
			visited_position[1] == obstacle_y {
			return
		}
	}
	new_tilemap.tiles[obstacle_y][obstacle_x] = true
	if new_tilemap.checkGuardLoop(tilemap.guard_x, tilemap.guard_y) {
		logrus.Debugf("%d %d\n", obstacle_x, obstacle_y)
		tilemap.tiles_obstructed[obstacle_y][obstacle_x] = true
	}
}

func (tilemap Map) checkGuardLoop(start_x int, start_y int) bool {
	// memorize the first step
	tilemap.guard_memory = append(tilemap.guard_memory, [3]int{tilemap.guard_x, tilemap.guard_y, tilemap.guard_direction})
	for tilemap.simulateGuardObstructedStep() {
		tilemap.print()
		logrus.Debugf("%d %d\n", tilemap.guard_memory[len(tilemap.guard_memory)-1][0], tilemap.guard_memory[len(tilemap.guard_memory)-1][1])
		logrus.Debugf("\n--------------\n")
		// if revisited the start position check if the guard arrived from the same direction
		for _, memory := range tilemap.guard_memory {
			if memory[0] == tilemap.guard_x &&
				memory[1] == tilemap.guard_y &&
				memory[2] == tilemap.guard_direction {
				return true
			}
		}
		// if len(tilemap.guard_memory) >= 2 &&
		// 	tilemap.guard_x == start_x &&
		// 	tilemap.guard_y == start_y {
		// 	logrus.Debugf("at the start again\n%d %d\n", tilemap.guard_x, tilemap.guard_y)
		// 	if tilemap.guard_memory[len(tilemap.guard_memory)-1][0] == tilemap.guard_x &&
		// 		tilemap.guard_memory[len(tilemap.guard_memory)-1][1] == tilemap.guard_y {
		// 		logrus.Debugf("looped!!!\n")
		// 		return true
		// 	}
		// }
		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			time.Sleep(500 * time.Millisecond)
		}
		// memorize every step
		tilemap.guard_memory = append(tilemap.guard_memory, [3]int{tilemap.guard_x, tilemap.guard_y, tilemap.guard_direction})
	}
	return false
}

func (tilemap *Map) simulateGuardObstructedStep() bool {
	// get next tile
	guard_x, guard_y, guard_left := tilemap.getNextPosition()

	// if guard left the field, return false
	if guard_left {
		return false
	}

	// if len(tilemap.guard_memory) == 0 || !(tilemap.guard_x == tilemap.guard_memory[len(tilemap.guard_memory)-1][0] && tilemap.guard_y == tilemap.guard_memory[len(tilemap.guard_memory)-1][1]) {
	// 	tilemap.guard_memory = append(tilemap.guard_memory, [2]int{tilemap.guard_x, tilemap.guard_y})
	// }
	// if guard hits a obstacle, rotate
	if tilemap.tiles[guard_y][guard_x] {
		tilemap.guard_direction = (tilemap.guard_direction + 1) % 4

		// mark the current tile as visited
		//tilemap.tiles_visited[tilemap.guard_y][tilemap.guard_x] = true
		// otherwise just walk
	} else {
		tilemap.guard_x = guard_x
		tilemap.guard_y = guard_y
	}
	return true
}
