package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) [][]string {
	antenna := [][]string{}

	file, err := os.Open(filename)

	if err != nil {
		return antenna
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sep := strings.Split(line, " ")
		aim := sep[0]
		sep[0] = aim[:len(aim)-1]

		antenna = append(antenna, sep)
	}
	return antenna
}

func main() {
	fmt.Print(readInput("day8.txt"))
}
