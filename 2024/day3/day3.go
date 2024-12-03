package day3

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(file string) int {
	lines := utils.ReadFileByLine(file)
	content := strings.Join(lines[:], ":")

	re_mul := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	multiplications := re_mul.FindAllString(content, -1)

	re_num := regexp.MustCompile(`[0-9]+`)
	result := 0
	for _, mul := range multiplications {
		numbers := re_num.FindAllString(mul, -1)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		result += (num1 * num2)
	}
	fmt.Println(result)
	return result
}

func Part2(file string) int {
	lines := utils.ReadFileByLine(file)
	content := strings.Join(lines[:], ":")

	re_mul := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don\'t\(\)`)
	inputs := re_mul.FindAllString(content, -1)

	do := true
	var multiplications []string
	for _, input := range inputs {
		if input == "don't()" {
			do = false
			continue
		} else if input == "do()" {
			do = true
			continue
		}

		if do {
			multiplications = append(multiplications, input)
		}
	}

	re_num := regexp.MustCompile(`[0-9]+`)
	result := 0
	for _, mul := range multiplications {
		numbers := re_num.FindAllString(mul, -1)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		result += (num1 * num2)
	}
	return result
}
