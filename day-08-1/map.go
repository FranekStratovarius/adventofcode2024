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
					logrus.Debugf("> %+v | %+v\n", antenna, second_antenna)
					position := Position{
						x: antenna.position.x*2 - second_antenna.position.x,
						y: antenna.position.y*2 - second_antenna.position.y,
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
