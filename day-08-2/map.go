package main

import (
	"github.com/sirupsen/logrus"
)

type Map struct {
	frequencies [][]rune
	antinodes   [][]bool
	antennas    map[rune][]Antenna
}

func (cityMap *Map) addRow(row []rune) {
	cityMap.frequencies = append(cityMap.frequencies, row)
	cityMap.antinodes = append(cityMap.antinodes, make([]bool, len(row)))

	if cityMap.antennas == nil {
		cityMap.antennas = make(map[rune][]Antenna)
	}

	for x, frequency := range row {
		if frequency != '.' {
			cityMap.antennas[frequency] = append(cityMap.antennas[frequency], Antenna{
				position: Position{
					x: x,
					y: len(cityMap.frequencies) - 1,
				},
				frequency: frequency,
			})
		}
	}
}

func (cityMap Map) print() {
	for y, row := range cityMap.frequencies {
		for x, frequency := range row {
			if cityMap.antinodes[y][x] {
				logrus.Debugf("#")
			} else {
				logrus.Debugf("%c", frequency)
			}
		}
		logrus.Debugf("\n")
	}
}

func (cityMap Map) createAntinodeMap() {
	for frequency, antennas := range cityMap.antennas {
		logrus.Debugf("%c\t%+v\n", frequency, antennas)
		for i, antenna := range antennas {
			for j, second_antenna := range antennas {
				if i != j {
					inside_map := true
					for k := 1; inside_map; k++ {
						// the tower itself also is an antinodw when in line with another tower
						cityMap.antinodes[antenna.position.y][antenna.position.x] = true
						logrus.Debugf("> [%d] %+v | %+v\n", k, antenna, second_antenna)
						logrus.Debugf(
							"difference: %d %d\n",
							antenna.position.x-second_antenna.position.x,
							antenna.position.y-second_antenna.position.y,
						)
						position := Position{
							x: antenna.position.x*(1+k) - second_antenna.position.x*k,
							y: antenna.position.y*(1+k) - second_antenna.position.y*k,
						}
						logrus.Debugf(
							"%d %d\n",
							position.x,
							position.y,
						)
						if position.x >= 0 &&
							position.y >= 0 &&
							position.x < len(cityMap.frequencies[0]) &&
							position.y < len(cityMap.frequencies) {
							cityMap.antinodes[position.y][position.x] = true
							logrus.Debugf("added pos\n")
						} else {
							inside_map = false
						}
					}
				}
			}
		}
	}
}

func (cityMap Map) countAntinodes() (sum int) {
	for _, row := range cityMap.antinodes {
		for _, antinode := range row {
			if antinode {
				sum++
			}
		}
	}
	return
}
