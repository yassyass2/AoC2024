package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

func CheckSum(aim int, products []int) bool{
	return true
}

func calculateResult(calibs [][]string) int {
	total := 0
	for _, val := range calibs{
		aim := val[0]
		val[0] = aim[:len(aim)-1]

		steps := make([]int, len(val))

		for i, str := range val {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting %q to int: %v\n", str, err)
				return 0
			}
			steps[i] = num
		}

		if CheckSum()()
	}
	return total
}

func main() {
	fmt.Print(readInput("day7.txt"))
}
