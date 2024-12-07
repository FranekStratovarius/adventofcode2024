package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type PlainFormatter struct{}

func (formatter *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf(entry.Message)), nil
}

func main() {
	logrus.SetFormatter(&PlainFormatter{})
	// logrus.SetLevel(logrus.DebugLevel)
	scanner := bufio.NewScanner(os.Stdin)

	tilemap := Map{}

	for scanner.Scan() {
		line := scanner.Text()
		// logrus.Printf("> %s\n", line)
		tilemap.addRow(line)
	}

	counter := 0
	tilemap.print()
	logrus.Debugf("\n--------------\n")
	tilemap.checkObstacle()
	logrus.Debugf("\n--------------\n")

	// tilemap.print()

	for tilemap.simulateStep() {
		counter++
		fmt.Printf("%d, ", counter)
		logrus.Debugf("\n%d --------------\n", counter)
		tilemap.checkObstacle()
		// 	time.Sleep(1000 * time.Millisecond)
	}
	fmt.Print("\n")

	logrus.Debugf("\n--------------\n")
	tilemap.print()
	logrus.Debugf("\n--------------\n")

	obstructed_tiles := tilemap.countObstructed()

	fmt.Printf("possible obstructions %d\n", obstructed_tiles)
}

/*
.........#
..........
..#.......
.......#..
..........
.#.O......
......OO#.
#O.O...V..
......#O..
*/
