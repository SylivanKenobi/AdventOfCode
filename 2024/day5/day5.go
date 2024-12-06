package day5

import (
	"adventofcode/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Part1(file string) int {
	lines := utils.ReadFileByLine(file)

	result := 0
	rules, updates := splitInput(lines)
	var validUpdates []int
	for _, update := range updates {
		valid := true
		for _, rule := range rules {
			i := slices.Index(update, rule[0])
			j := slices.Index(update, rule[1])
			if i >= 0 && j >= 0 && i > j {
				valid = false
			}
		}
		if valid {
			middle := update[len(update)/2]
			validUpdates = append(validUpdates, middle)
		}
	}
	for _, validUpdate := range validUpdates {
		result += validUpdate
	}
	return result
}

func Part2(file string) int {
	lines := utils.ReadFileByLine(file)

	result := 0
	rules, updates := splitInput(lines)
	var validUpdates []int
	for _, update := range updates {
		inValid := false
		pass := false
		for !pass {
			pass = true
			for _, rule := range rules {
				i := slices.Index(update, rule[0])
				j := slices.Index(update, rule[1])
				if i >= 0 && j >= 0 && i > j {
					tmpi := update[i]
					update[i] = update[j]
					update[j] = tmpi
					inValid = true
					pass = false
				}
			}
		}
		if inValid {
			middle := update[len(update)/2]
			validUpdates = append(validUpdates, middle)
		}
	}
	for _, validUpdate := range validUpdates {
		result += validUpdate
	}
	return result
}

func splitInput(input []string) ([][]int, [][]int) {
	blank := false
	var rules [][]int
	var updates [][]int
	for _, line := range input {
		if line != "" {
			if !blank {
				rule := strToIntArr(regexp.MustCompile(`[0-9]*`).FindAllString(line, -1))
				rules = append(rules, rule)
			} else {
				updates = append(updates, strToIntArr(strings.Split(line, ",")))
			}
		} else {
			blank = true
		}
	}
	return rules, updates
}

func strToIntArr(input []string) []int {
	var r []int
	for _, n := range input {
		a, _ := strconv.Atoi(n)
		r = append(r, a)
	}
	return r
}
