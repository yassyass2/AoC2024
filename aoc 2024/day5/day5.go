package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([][]string, [][]string) {
	rules := [][]string{}
	sequence := [][]string{}

	file, err := os.Open(filename)

	if err != nil {
		return rules, sequence
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, "|") {
			rules = append(rules, strings.Split(line, "|"))
		} else if strings.Contains(line, ",") {
			sequence = append(sequence, strings.Split(line, ","))
		}
	}

	return rules, sequence
}

func AbidesByAllRules(update []string, rules [][]string) bool {
	firstIndex := -1
	secondIndex := -1
	for i, value := range update {
		if value == rules[0][0] && firstIndex == -1 {
			firstIndex = i
		}
		if value == rules[0][1] && secondIndex == -1 {
			secondIndex = i
		}
		if firstIndex != -1 && secondIndex != -1 {
			break
		}
	}
	if firstIndex == -1 || secondIndex == -1 || firstIndex < secondIndex {
		if len(rules) > 1 {
			return true && AbidesByAllRules(update, rules[1:])
		}
		return true
	}
	return false
}

func Swapper(update []string, rules [][]string) []string {
	firstIndex := -1
	secondIndex := -1
	for i, value := range update {
		if value == rules[0][0] && firstIndex == -1 {
			firstIndex = i
		}
		if value == rules[0][1] && secondIndex == -1 {
			secondIndex = i
		}
		if firstIndex != -1 && secondIndex != -1 {
			break
		}
	}
	if firstIndex != -1 && secondIndex != -1 && firstIndex > secondIndex {
		firstVal := update[firstIndex]
		update[firstIndex] = update[secondIndex]
		update[secondIndex] = firstVal
	}
	if len(rules) > 1 {
		return Swapper(update, rules[1:])
	}
	return update
}

func countMiddles(sequence [][]string, rules [][]string) []int {
	total := 0
	swapped := 0
	for _, CurrentUpdate := range sequence {

		if AbidesByAllRules(CurrentUpdate, rules) {
			ToAdd, err := strconv.Atoi(CurrentUpdate[(len(CurrentUpdate)-1)/2])
			if err != nil {
				fmt.Println("Error:", err)
			}
			total += ToAdd
		} else {
			for !AbidesByAllRules(CurrentUpdate, rules) {
				CurrentUpdate = Swapper(CurrentUpdate, rules)
			}
			ToSwap, err := strconv.Atoi(CurrentUpdate[(len(CurrentUpdate)-1)/2])
			if err != nil {
				fmt.Println("Error:", err)
			}
			swapped += ToSwap
		}
	}
	return []int{total, swapped}
}

func main() {
	rules, sequence := readInput("day5.txt")

	result := countMiddles(sequence, rules)
	fmt.Printf("%d", result[0])
	fmt.Printf(" ")
	fmt.Printf("%d", result[1])
}
