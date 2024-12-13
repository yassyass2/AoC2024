package main

import (
	"bufio"
	"fmt"
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

func findStart(m [][]string) []int {
	for i, lane := range m {
		for j, spot := range lane {
			if spot != "." && spot != "#" {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func move(m [][]string, p []int, mapping map[string][]int) [][]string {
	if p[0] < 1 || p[0] > 128 || p[1] < 1 || p[1] > 128 {
		return m
	}
	old := p
	symbol := m[p[0]][p[1]]
	change_vector := mapping[symbol]

	p[0] += change_vector[0]
	p[1] += change_vector[1]
	if p[0] < 1 || p[0] > 129 || p[1] < 1 || p[1] > 129 || m[p[0]][p[1]] == "#" {
		p[0], p[1] = old[0], old[1]
		if symbol == "^" {
			m[p[0]][p[1]] = ">"
		} else if symbol == ">" {
			m[p[0]][p[1]] = "v"
		} else if symbol == "v" {
			m[p[0]][p[1]] = "<"
		} else {
			m[p[0]][p[1]] = "^"
		}
	} else {
		m[old[0]][old[1]] = "X"
		if symbol == "^" {
			m[p[0]][p[1]] = ">"
		} else if symbol == ">" {
			m[p[0]][p[1]] = "v"
		} else if symbol == "v" {
			m[p[0]][p[1]] = "<"
		} else {
			m[p[0]][p[1]] = "^"
		}
	}
	newP := make([]int, len(p))
	newM := make([][]string, len(m))
	copy(newM, m)
	copy(newP, p)
	//fmt.Print(countPath(newM))
	return move(newM, newP, mapping)
}

func countPath(m [][]string) int {
	count := 0
	for _, lane := range m {
		for _, spot := range lane {
			if spot == "X" {
				count++
			}
		}
	}
	return count
}

func main() {
	map_to_patrol := readInput("day6.txt")

	mapping := make(map[string][]int)
	mapping["^"] = []int{-1, 0}
	mapping[">"] = []int{0, 1}
	mapping["v"] = []int{1, 0}
	mapping["<"] = []int{0, -1}

	map_after := move(map_to_patrol, findStart(map_to_patrol), mapping)
	fmt.Print(countPath(map_after))
	//fmt.Printf("%d")
	//fmt.Printf("\n")

}
