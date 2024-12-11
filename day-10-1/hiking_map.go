package main

import (
	"github.com/sirupsen/logrus"
)

type HikingMap struct {
	heightmap    [][]int
	hikingTrails []HikingTrail
}

type Position struct {
	x int
	y int
}

type HikingTrail struct {
	positions []Position
}

func (hikingMap *HikingMap) addRow(line string) {
	var row []int
	for _, tile := range line {
		row = append(row, int(tile-'0'))
	}
	hikingMap.heightmap = append(hikingMap.heightmap, row)
}

func (hikingMap HikingMap) print() {
	for _, row := range hikingMap.heightmap {
		for _, tile := range row {
			logrus.Debugf("%d", tile)
		}
		logrus.Debugf("\n")
	}
}

func (hikingMap *HikingMap) findHikingTrails() (trailheadRatingSum int) {
	for y, row := range hikingMap.heightmap {
		for x, tile := range row {
			if tile == 0 {
				var hikingTrails []HikingTrail
				var trailpeaks []Position
				hikingMap.findHikingTrailsRecursive(
					&hikingTrails,
					HikingTrail{},
					Position{
						x: x,
						y: y,
					},
				)
				// logrus.Debugf("\n")
				for _, trail := range hikingTrails {
					// logrus.Debugf("%+v\n", trail.positions)
					alreadyPresent := false
					for _, trailpeak := range trailpeaks {
						if trail.positions[9].x == trailpeak.x &&
							trail.positions[9].y == trailpeak.y {
							alreadyPresent = true
							break
						}
					}
					if !alreadyPresent {
						trailpeaks = append(trailpeaks, trail.positions[9])
					}
					// logrus.Debugf("%+v\n", trail.positions[9])
				}
				hikingMap.hikingTrails = append(hikingMap.hikingTrails, hikingTrails...)
				// logrus.Debugf("trailheads: %v\n", trailheads)
				trailheadRatingSum += len(trailpeaks)
			}
		}
	}
	return
}

func (hikingMap *HikingMap) findHikingTrailsRecursive(hikingTrails *[]HikingTrail, hikingTrail HikingTrail, newPosition Position) {
	// deep copy positions
	newPositions := make([]Position, len(hikingTrail.positions))
	copy(newPositions, hikingTrail.positions)
	newPositions = append(newPositions, newPosition)

	// set positions of hiking trail
	hikingTrail.positions = newPositions

	height := len(hikingTrail.positions) - 1

	// if position is outside of map, search has failed
	if newPosition.y < 0 ||
		newPosition.x < 0 ||
		newPosition.y >= len(hikingMap.heightmap) ||
		newPosition.x >= len(hikingMap.heightmap[0]) {
		// logrus.Debugf("outside\n")
		return
	}

	// stop if current height is not the desired height
	if hikingMap.heightmap[newPosition.y][newPosition.x] != height {
		return
	}

	// if height 9 is reached, search is successful
	if height == 9 {
		// fmt.Printf("finished: %+v\n", hikingTrail.positions)
		*hikingTrails = append(*hikingTrails, hikingTrail)
		return
	}

	// check in each direction
	positionUp := newPosition
	positionUp.y++
	// hikingMap.findHikingTrailsRecursive(hikingTrails, hikingTrail, Position{x: newPosition.x, y: newPosition.y + 1})
	hikingMap.findHikingTrailsRecursive(hikingTrails, hikingTrail, positionUp)
	positionRight := newPosition
	positionRight.x++
	hikingMap.findHikingTrailsRecursive(hikingTrails, hikingTrail, positionRight)
	positionDown := newPosition
	positionDown.y--
	hikingMap.findHikingTrailsRecursive(hikingTrails, hikingTrail, positionDown)
	positionLeft := newPosition
	positionLeft.x--
	hikingMap.findHikingTrailsRecursive(hikingTrails, hikingTrail, positionLeft)
}
