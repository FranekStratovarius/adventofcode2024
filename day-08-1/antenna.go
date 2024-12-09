package main

type Position struct {
	x int
	y int
}

type Antenna struct {
	position  Position
	frequency rune
}