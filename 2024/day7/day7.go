package day7

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
)

// Always turn right till stopped
func Part1(file string) int {
	lines := utils.ReadFileByLine(file)
	result := 0
	re := regexp.MustCompile(`[0-9]+`)
	for _, line := range lines {
		a := re.FindAllString(line, -1)
		test, _ := strconv.Atoi(a[0])
		numbers := utils.StrToIntArr(a[1:len(a)])
		if valid(numbers, test, 1, numbers[0], 1) {
			result += test
		}
	}

	return result
}

func Part2(file string) int {
	lines := utils.ReadFileByLine(file)
	result := 0
	re := regexp.MustCompile(`[0-9]+`)
	for _, line := range lines {
		a := re.FindAllString(line, -1)
		test, _ := strconv.Atoi(a[0])
		numbers := utils.StrToIntArr(a[1:len(a)])
		if valid(numbers, test, 1, numbers[0], 2) {
			result += test
		}
	}

	return result
}

func valid(numbers []int, test int, i int, res int, part int) bool {
	if i == len(numbers) {
		return test == res
	} else {
		if part == 1 {
			return valid(numbers, test, i+1, res+numbers[i], part) || valid(numbers, test, i+1, res*numbers[i], part)
		} else {
			concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", res, numbers[i]))
			return valid(numbers, test, i+1, res+numbers[i], part) || valid(numbers, test, i+1, res*numbers[i], part) || valid(numbers, test, i+1, concat, part)
		}
	}

}
