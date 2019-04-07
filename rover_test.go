package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestMoveRover(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 1 N ", "1 1 N"},
		{"1 1 N M", "1 2 N"},
		{"1 1 E M", "2 1 E"},
		{"1 1 W M", "0 1 W"},
		{"1 1 S M", "1 0 S"},
		{"0 0 S M", "0 0 S"},
		{"5 5 N M", "5 5 N"},
		{"1 2 N LMLMLMLMM", "1 3 N"},
		{"3 3 E MMRMMRMRRM", "5 1 E"},
	}
	plateau := Plateau{x: 5, y: 5}
	for _, tt := range tests {
		tmp := strings.Split(tt.input, " ")
		x, _ := strconv.Atoi(tmp[0])
		y, _ := strconv.Atoi(tmp[1])
		heading := tmp[2]
		steps := tmp[3]
		rover := NewRover(x, y, heading)
		rover.ExecuteIntructions([]rune(steps), plateau)

		if rover.String() != tt.expected {
			t.Errorf("rover got wrong position. got=%v, want=%v", rover.String(), tt.expected)
		}
	}
}
