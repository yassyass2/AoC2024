package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		sep := strings.Split(line, " ")
		aim := sep[0]
		sep[0] = aim[:len(aim)-1]

		calibrations = append(calibrations, sep)
	}
	return calibrations
}

func CheckSum(aim int, accum int, products []int) bool {
	if len(products) == 1 {
		return accum+products[0] == aim || accum*products[0] == aim
	}
	return CheckSum(aim, accum*products[0], products[1:]) || CheckSum(aim, accum+products[0], products[1:])
}

func calculateResult(calibs [][]string) int {
	total := 0
	for _, val := range calibs {
		steps := make([]int, len(val))

		for i, str := range val {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting %q to int: %v\n", str, err)
				return 0
			}
			steps[i] = num
		}

		if CheckSum(steps[0], steps[1], steps[2:]) {
			total += steps[0]
		}
	}
	return total
}

func main() {
	fmt.Print(calculateResult(readInput("day7.txt")))
}