package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) [][]string {
	calibrations := [][]string{}

	file, err := os.Open(filename)

	if err != nil {
		return calibrations
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		calibrations = append(calibrations, strings.Split(line, " "))
	}
	return calibrations
}

func calculateResult(calibs [][]string) int {
	total := 0
	return total
}

func main() {
	fmt.Print(readInput("day7.txt"))
}
