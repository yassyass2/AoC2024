package main

import (
	"bufio"
	"os"
	"strings"
)

func readInput(filename string) [][]string {
	patrol := [][]string{}

	file, err := os.Open(filename)

	if err != nil {
		return patrol
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		patrol = append(patrol, strings.Split(line, ""))
	}
	return patrol
}

func main() {
	map_to_patrol := readInput("day6.txt")

	//fmt.Printf("%d")
	//fmt.Printf("\n")

}
