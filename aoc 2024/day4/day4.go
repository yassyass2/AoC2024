package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countOccurrences(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)
	count := 0

	wordRunes := []rune(word)

	checkDirection := func(r, c, dr, dc int) bool {
		for i := 0; i < wordLen; i++ {
			rr := r + dr*i
			cc := c + dc*i
			if rr < 0 || rr >= rows || cc < 0 || cc >= cols || grid[rr][cc] != wordRunes[i] {
				return false
			}
		}
		return true
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			directions := [][2]int{
				{0, 1}, {1, 0}, {1, 1}, {1, -1},
				{0, -1}, {-1, 0}, {-1, -1}, {-1, 1},
			}
			for _, dir := range directions {
				if checkDirection(r, c, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func countOccurrencesP2(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	//wordLen := len(word)
	count := 0

	wordRunes := []rune(word)

	checkDirection := func(y int, x int, directions [][2]int) bool {

		linksbovenX := x + directions[2][0]
		linksbovenY := y + directions[2][1]
		linksonderX := x + directions[3][0]
		linksonderY := y + directions[3][1]

		rechtsbovenX := x + directions[1][0]
		rechtsbovenY := y + directions[1][1]
		rechtsonderX := x + directions[0][0]
		rechtsonderY := y + directions[0][1]

		cords := []int{linksbovenX, linksbovenY, linksonderX, linksonderY,
			rechtsbovenX, rechtsbovenY, rechtsonderX, rechtsonderY}

		for _, col := range cords {
			if col < 0 {
				return false
			}
		}
		if linksbovenX >= cols || linksonderX >= cols || rechtsbovenX >= cols || rechtsonderX >= cols ||
			linksbovenY >= rows || linksonderY >= rows || rechtsbovenY >= rows || rechtsonderY >= rows {
			return false
		}

		return grid[linksbovenY][linksbovenX] == wordRunes[0] && grid[rechtsonderY][rechtsonderX] == wordRunes[2] && grid[linksonderY][linksonderX] == wordRunes[0] && grid[rechtsbovenY][rechtsbovenX] == wordRunes[2] ||
			grid[linksbovenY][linksbovenX] == wordRunes[2] && grid[rechtsonderY][rechtsonderX] == wordRunes[0] && grid[linksonderY][linksonderX] == wordRunes[0] && grid[rechtsbovenY][rechtsbovenX] == wordRunes[2] ||
			grid[linksbovenY][linksbovenX] == wordRunes[0] && grid[rechtsonderY][rechtsonderX] == wordRunes[2] && grid[linksonderY][linksonderX] == wordRunes[2] && grid[rechtsbovenY][rechtsbovenX] == wordRunes[0] ||
			grid[linksbovenY][linksbovenX] == wordRunes[2] && grid[rechtsonderY][rechtsonderX] == wordRunes[0] && grid[linksonderY][linksonderX] == wordRunes[2] && grid[rechtsbovenY][rechtsbovenX] == wordRunes[0]
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			directions := [][2]int{
				{1, 1}, {1, -1},
				{-1, -1}, {-1, 1},
			}
			if grid[r][c] == wordRunes[1] {
				if checkDirection(r, c, directions) {
					count++
				}
			}
		}
	}

	return count
}

func readGridFromFile(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func main() {
	filename := "day4.txt"

	grid, err := readGridFromFile(filename)
	if err != nil {
		fmt.Println("Error reading grid:", err)
		return
	}

	count := countOccurrences(grid, "XMAS")
	fmt.Printf("%d", count)

	fmt.Printf("  ")

	count2 := countOccurrencesP2(grid, "MAS")
	fmt.Printf("%d", count2)

}
