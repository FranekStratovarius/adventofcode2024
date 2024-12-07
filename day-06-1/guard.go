package main

func (tilemap *Map) simulateStep() bool {
	guard_x := tilemap.guard_x
	guard_y := tilemap.guard_y
	// mark the current tile as visited
	tilemap.tiles_visited[guard_y][guard_x] = true
	// get next tile
	switch tilemap.guard_direction {
	case 0:
		guard_y--
	case 1:
		guard_x++
	case 2:
		guard_y++
	case 3:
		guard_x--
	}
	// if guard left the field, return false
	if guard_x < 0 || guard_y < 0 ||
		guard_x > len(tilemap.tiles[0])-1 ||
		guard_y > len(tilemap.tiles)-1 {
		return false
	}

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
