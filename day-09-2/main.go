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

	diskLine := ""

	for scanner.Scan() {
		line := scanner.Text()
		logrus.Debugf("> %s\n", line)
		diskLine = line
	}

	disk := readDisk(diskLine)
	// disk.print()
	disk.fragment()

	checksum := disk.calculateChecksum()

	fmt.Printf("the checksum is: %d\n", checksum)
}
