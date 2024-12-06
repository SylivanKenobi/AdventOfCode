package day6

import (
	"adventofcode/utils"
	"regexp"
	"strconv"
)

// Always turn right till stopped
func Part1(file string) int {
	lines := utils.ReadFileByLine(file)

	result := 0
	upRe := regexp.MustCompile(`^`)
	dwRe := regexp.MustCompile(`v`)
	riRe := regexp.MustCompile(`>`)
	leRe := regexp.MustCompile(`<`)
	for _, line := range lines {

	}
	return result
}

func Part2(file string) int {
	lines := utils.ReadFileByLine(file)

	result := 0

	return result
}

func strToIntArr(input []string) []int {
	var r []int
	for _, n := range input {
		a, _ := strconv.Atoi(n)
		r = append(r, a)
	}
	return r
}
