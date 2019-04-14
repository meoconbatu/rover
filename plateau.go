package main

// Plateau the upper-right coordinates of the plateau
type Plateau struct {
	x, y int
}

// NewPlateau create a new plateau
func NewPlateau(x, y int) Plateau {
	return Plateau{x: x, y: y}
}
