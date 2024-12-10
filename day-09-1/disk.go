package main

import "github.com/sirupsen/logrus"

type Segment struct {
	free bool
	id   int
}
type Disk struct {
	segments []Segment
}

func (disk Disk) print() {
	for _, segment := range disk.segments {
		if !segment.free {
			logrus.Debugf("%d", segment.id)
		} else {
			logrus.Debugf(".")
		}
	}
	logrus.Debugf("\n")
}

func readDisk(diskLine string) (disk Disk) {
	counter := 0
	for i, segment := range diskLine {
		free := i%2 == 1
		for i := 0; i < int(segment-'0'); i++ {
			disk.segments = append(disk.segments, Segment{
				free: free,
				id:   counter,
			})
		}
		if !free {
			counter++
		}
	}
	return
}

func (disk *Disk) fragment() {
	disk.print()
	// don'tdo anything if no freespot is available
	if !disk.checkEmptySpotAvailable() {
		return
	}
	left_counter := 0
	for right_counter := len(disk.segments) - 1; right_counter > left_counter; right_counter-- {
		if !disk.segments[right_counter].free {
			// continue searching from last time untila free spot is found
			left_counter = disk.findFreeSpot(left_counter)
			if right_counter > left_counter {
				logrus.Debugf("swapping: %d -> %d\n", right_counter, left_counter)
				disk.segments[left_counter].free = false
				disk.segments[left_counter].id = disk.segments[right_counter].id
				disk.segments[right_counter].free = true
				// disk.print()
			}
		}
	}
}

func (disk Disk) checkEmptySpotAvailable() bool {
	for i := 0; i < len(disk.segments); i++ {
		if disk.segments[i].free {
			return true
		}
	}
	return false
}

func (disk Disk) findFreeSpot(start int) (freeSpot int) {
	for i := start; i < len(disk.segments); i++ {
		if disk.segments[i].free {
			freeSpot = i
			return
		}
	}
	return
}

func (disk Disk) calculateChecksum() (checksum int) {
	for i, segment := range disk.segments {
		if !segment.free {
			checksum += i * segment.id
		}
	}
	return
}
