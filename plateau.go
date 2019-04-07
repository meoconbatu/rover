package main

type Plateau struct {
	x, y int
}

func NewPlateau(x, y int) Plateau {
	return Plateau{x: x, y: y}
}
