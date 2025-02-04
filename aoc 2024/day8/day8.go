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

func getMappedNodes(Map [][]string) map[string][][]int {
	MappedNodes := make(map[string][][]int)
	for y, ant := range Map {
		for x, node := range ant {
			if node != "." {
				if _, exists := MappedNodes[node]; exists {
					MappedNodes[node] = append(MappedNodes[node], []int{y, x})
				} else {
					MappedNodes[node] = [][]int{{y, x}}
				}
			}
		}
	}
	return MappedNodes
}

func main() {
	fmt.Print(getMappedNodes(readInput("day8.txt")))
}
