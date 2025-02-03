package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) [][]string {
	Map := [][]string{}

	file, err := os.Open(filename)

	if err != nil {
		return Map
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sep := strings.Split(line, " ")
		aim := sep[0]
		sep[0] = aim[:len(aim)-1]

		Map = append(Map, sep)
	}
	return Map
}

func main() {
	fmt.Print(readInput("day8.txt"))
}
