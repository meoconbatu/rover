package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var inputFileName, outputFileName string
	var runCmd = &cobra.Command{
		Use:   "run [string to echo]",
		Short: "Run rover",
		Long: `echo things multiple times back to the user by providing
	a count and a string.`,
		Run: func(cmd *cobra.Command, args []string) {
			Run(inputFileName, outputFileName)
		},
	}

	runCmd.Flags().StringVarP(&inputFileName, "input", "i", "", "input file name")
	runCmd.MarkFlagRequired("input")

	runCmd.Flags().StringVarP(&outputFileName, "output", "o", "", "output file name")

	var rootCmd = &cobra.Command{Use: "rover"}
	rootCmd.AddCommand(runCmd)
	rootCmd.Execute()
}
func Run(inputFileName, outputFileName string) {
	var plateau *Plateau
	rovers, plateau, err := readInputFromFile(inputFileName)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	for i := 0; i < len(rovers); i++ {
		rovers[i].ExecuteInstructions(*plateau)
	}
	if outputFileName != "" {
		writeOutputToFile(rovers, outputFileName)
	} else {
		writeOutputToConsole(rovers)
	}
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
func writeOutputToFile(rovers []Rover, outputFileName string) error {
	f, err := os.Create(outputFileName)
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
