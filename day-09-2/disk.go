package main

import (
	"github.com/sirupsen/logrus"
)

type Segment struct {
	free   bool
	id     int
	length int
	prev   *Segment
	next   *Segment
}
type Disk struct {
	head *Segment
	tail *Segment
}

func (disk Disk) print() {
	currentSegment := disk.head
	for currentSegment != nil {
		for i := 0; i < currentSegment.length; i++ {
			if !currentSegment.free {
				logrus.Debugf("%d", currentSegment.id)
			} else {
				logrus.Debugf(".")
			}
		}
		currentSegment = currentSegment.next
	}
	logrus.Debugf("\n")
}

func readDisk(diskLine string) (disk Disk) {
	counter := 0
	currentSegment := disk.head
	for i, segment := range diskLine {
		free := i%2 == 1
		nextSegment := &Segment{
			free:   free,
			id:     counter,
			length: int(segment - '0'),
			prev:   currentSegment,
		}
		if currentSegment == nil {
			disk.head = nextSegment
		} else {
			currentSegment.next = nextSegment
		}
		if !free {
			counter++
		}
		currentSegment = nextSegment
		disk.tail = currentSegment
	}
	return
}

func (disk *Disk) fragment() {
	disk.print()
	// don'tdo anything if no free spot is available
	if !disk.checkEmptySpotAvailable() {
		return
	}

	currentSegment := disk.tail
	for currentSegment != nil {
		if !currentSegment.free {
			freeSegment, available := disk.findFreeSpot(currentSegment)
			if available {
				swap(freeSegment, currentSegment)
				disk.print()
			}
		}
		currentSegment = currentSegment.prev
	}
}

func (disk Disk) checkEmptySpotAvailable() bool {
	currentSegment := disk.head
	for currentSegment != nil {
		if currentSegment.free {
			return true
		}
		currentSegment = currentSegment.next
	}
	return false
}

func (disk Disk) findFreeSpot(file *Segment) (freeSegment *Segment, available bool) {
	currentSegment := disk.head
	for currentSegment != nil && currentSegment != file {
		if currentSegment.free && currentSegment.length >= file.length {
			freeSegment = currentSegment
			available = true
			return
		}
		currentSegment = currentSegment.next
	}

	available = false
	return
}

func swap(freeSegment *Segment, file *Segment) {
	// save length of free spot
	freeSpotLength := freeSegment.length
	// write file id and length into free spot
	freeSegment.free = false
	freeSegment.id = file.id
	freeSegment.length = file.length
	// if free spot is left, create new free spot after file
	if freeSpotLength > file.length {
		newFreeSegment := &Segment{
			free:   true,
			length: freeSpotLength - file.length,
			prev:   freeSegment,
			next:   freeSegment.next,
		}
		freeSegment.next.prev = newFreeSegment
		freeSegment.next = newFreeSegment
	}

	// write free spot into old file
	file.free = true
	// merge free spot with potential free spots before and after
	if file.prev != nil && file.prev.free {
		// save pointer
		otherFreeSegment := file.prev
		// add length of previous segment to current segment
		file.length += otherFreeSegment.length
		// change pointer to previous segment to current segment
		otherFreeSegment.prev.next = file
		// change pointer to previous segment to segment before previous
		file.prev = otherFreeSegment.prev
	}
	if file.next != nil && file.next.free {
		// save pointer
		otherFreeSegment := file.next
		// add length of previous segment to current segment
		file.length += otherFreeSegment.length
		// change pointer to next segment to current segment
		if otherFreeSegment.next != nil {
			otherFreeSegment.next.prev = file
		}
		// change pointer to next segment to segment after next
		file.next = otherFreeSegment.next
	}
}

func (disk Disk) calculateChecksum() (checksum int) {
	currentSegment := disk.head
	counter := 0
	for currentSegment != nil {
		for i := 0; i < currentSegment.length; i++ {
			if !currentSegment.free {
				checksum += counter * currentSegment.id
			}
			counter++
		}
		currentSegment = currentSegment.next
	}
	return
}

/*

func (disk Disk) findFreeSpot(length int) (freeSpot int, available bool) {
	for i := 0; i < len(disk.segments); i++ {
		if disk.segments[i].free && disk.segments[i].length >= length {
			freeSpot = i
			available = true
			return
		}
	}
	available = false
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

func (disk *Disk) swap(freeSpot int, file int) {
	// save length of free spot
	freeSpotLength := disk.segments[freeSpot].length
	// write file id and length into free spot
	disk.segments[freeSpot].free = false
	disk.segments[freeSpot].id = disk.segments[file].id
	disk.segments[freeSpot].length = disk.segments[file].length
	// if free spot is left, create new free spot after file
	if freeSpotLength > disk.segments[file].length {
		disk.segments = slices.Insert(disk.segments, freeSpot+1, Segment{
			free:   true,
			length: freeSpotLength - disk.segments[file].length,
		})
		file++
	}

	// write free spot into old file
	disk.segments[file].free = true
	// merge free spot with potential free spots before and after
}

*/
