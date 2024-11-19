package day1_part1

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
)

// Executes the puzzle for all days
func Part1(file string) int {
	lines := utils.ReadFileByLine("1_1.txt")
	calibration := 0
	for _, line := range lines {
		rgx := regexp.MustCompile(`\d`)
		numbers := rgx.FindAllString(line, -1)
		coordinate, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		calibration += coordinate

	}
	fmt.Println(calibration)
	return calibration
}
