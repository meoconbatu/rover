package main

import (
	"bytes"
	"strconv"
)

type Rover struct {
	x, y    int
	heading string
	// steps   []rune
}

func NewRover(x, y int, heading string) Rover {
	return Rover{x: x, y: y, heading: heading}
}
func (r Rover) String() string {
	var out bytes.Buffer
	out.WriteString(strconv.Itoa(r.x) + " " + strconv.Itoa(r.y) + " " + r.heading)
	return out.String()
}
func (r *Rover) ExecuteIntructions(instructions []rune) {
	for _, instruction := range instructions {
		switch instruction {
		case 'M':
			r.move()
		case 'L':
			r.turnLeft()
		case 'R':
			r.turnRight()
		}
	}
}
func (r *Rover) move() {
	switch r.heading {
	case "N":
		r.y = r.y + 1
	case "E":
		r.x = r.x + 1
	case "W":
		r.x = r.x - 1
	case "S":
		r.y = r.y - 1
	}
	if r.x < 0 {
		r.x = 0
	}
	if r.y < 0 {
		r.y = 0
	}
}
func (r *Rover) turnLeft() {
	switch r.heading {
	case "N":
		r.heading = "W"
	case "E":
		r.heading = "N"
	case "W":
		r.heading = "S"
	case "S":
		r.heading = "E"
	}
}
func (r *Rover) turnRight() {
	switch r.heading {
	case "N":
		r.heading = "E"
	case "E":
		r.heading = "S"
	case "W":
		r.heading = "N"
	case "S":
		r.heading = "W"
	}
}
