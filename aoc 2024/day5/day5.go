package main

import (
	"bufio"
	"os"
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

func countMiddles(sequence [][]string, rules [][]string) {
	total := 0
	for i, instance := range sequence {

		for j, value := range instance {

		}
	}
}
