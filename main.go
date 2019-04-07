package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	xmax, ymax int
)

func main() {
	rovers, err := ReadInputFromFile("input.txt")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	err = WriteOutputToFile(rovers)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}
func ReadInputFromFile(fileName string) ([]Rover, error) {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}

	inputLines := strings.Split(string(input), "\n")

	tmp := strings.Split(inputLines[0], " ")
	xmax, _ = strconv.Atoi(tmp[0])
	ymax, _ = strconv.Atoi(tmp[1])
	plateau := NewPlateau(xmax, ymax)
	rovers := make([]Rover, xmax*ymax)
	for i := 1; i < len(inputLines); i = i + 2 {
		tmp := strings.Split(inputLines[i], " ")
		x, _ := strconv.Atoi(tmp[0])
		y, _ := strconv.Atoi(tmp[1])
		heading := tmp[2]
		instructions := inputLines[i+1]

		rover := NewRover(x, y, heading)
		rover.ExecuteIntructions([]rune(instructions), plateau)
		rovers = append(rovers, rover)
	}
	return rovers, nil
}
func WriteOutputToFile(rovers []Rover) error {
	f, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, rover := range rovers {
		_, err = w.WriteString(rover.String() + "\n")
		if err != nil {
			return err
		}
	}
	w.Flush()
	return nil
}
