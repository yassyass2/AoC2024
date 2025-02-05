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

func specificNode(cords [][]int) [][]int {
	resultSet := [][]int{}
	for index, cord := range cords {
		for ind, coor := range cords {
			if index == ind {
				continue
			}
			underY := cord[0] - (cord[0] - coor[0])
			underX := cord[1] - (cord[1] - coor[1])
			upperY := cord[0] + (cord[0] - coor[0])
			upperX := cord[1] + (cord[1] - coor[1])
			if underY >= 0 && underY <= 49 && underX >= 0 && underX <= 49 {
				resultSet = append(resultSet, []int{underY, underX})
			}
			if upperY >= 0 && upperY <= 49 && upperX >= 0 && upperX <= 49 {
				resultSet = append(resultSet, []int{upperY, upperX})
			}
		}
	}
	return resultSet
}

func main() {
	original := readInput("day8.txt")
	mapped := getMappedNodes(original)
	fmt.Print(mapped)
	for _, value := range mapped {
		specificNode(value)
	}
}
