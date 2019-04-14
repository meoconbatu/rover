package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var plateau *Plateau
	if len(os.Args) == 1 {
		fmt.Println("missing input file name")
		return
	}
	inputFileName := os.Args[1]

	rovers, plateau, err := readInputFromFile(inputFileName)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	for i := 0; i < len(rovers); i++ {
		rovers[i].ExecuteInstructions(*plateau)
	}
	// err = writeOutputToFile(rovers)
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return
	// }
	writeOutputToConsole(rovers)
}
func readInputFromFile(fileName string) ([]Rover, *Plateau, error) {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, nil, err
	}

	inputLines := strings.Split(string(input), "\n")

	tmp := strings.Split(inputLines[0], " ")
	xmax, _ := strconv.Atoi(tmp[0])
	ymax, _ := strconv.Atoi(tmp[1])
	plateau := NewPlateau(xmax, ymax)

	rovers := make([]Rover, 0)

	for i := 1; i < len(inputLines); i = i + 2 {
		tmp := strings.Split(inputLines[i], " ")
		x, _ := strconv.Atoi(tmp[0])
		y, _ := strconv.Atoi(tmp[1])
		heading := tmp[2]
		instructions := inputLines[i+1]
		rover := NewRover(x, y, heading, []rune(instructions))
		rovers = append(rovers, rover)
	}
	return rovers, &plateau, nil
}
func writeOutputToFile(rovers []Rover) error {
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
func writeOutputToConsole(rovers []Rover) {
	for _, rover := range rovers {
		fmt.Println(rover.String())
	}
}
