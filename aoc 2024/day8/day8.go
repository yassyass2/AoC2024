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
		sep := strings.Split(line, "")
		Map = append(Map, sep)
	}
	return Map
}

func getNodes(Map [][]string) map[string][][]int {
	SeperatedNodes := make(map[string][][]int)
	for y, ant := range Map {
		for x, node := range ant {
			if node != "." {
				if _, exists := SeperatedNodes[node]; exists {
					SeperatedNodes[node] = append(SeperatedNodes[node], []int{y, x})
				} else {
					SeperatedNodes[node] = [][]int{{y, x}}
				}
			}
		}
	}
	return SeperatedNodes
}

func main() {
	fmt.Print(getNodes(readInput("day8.txt")))
}
