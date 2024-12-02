package day2

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(file string) int {
	lines := utils.ReadFileByLine(file)
	validLine := 0
	for _, line := range lines {
		sNumbers := strings.Split(line, " ")
		var numbers []int
		for _, sn := range sNumbers {
			num, _ := strconv.Atoi(sn)
			numbers = append(numbers, num)
		}
		if Validline(numbers) {
			validLine += 1
		}
	}
	fmt.Println(validLine)
	return validLine
}

func Part2(file string) int {
	lines := utils.ReadFileByLine(file)
	validLine := 0
	for _, line := range lines {
		sNumbers := strings.Split(line, " ")
		var numbers []int
		for _, sn := range sNumbers {
			num, _ := strconv.Atoi(sn)
			numbers = append(numbers, num)
		}
		valid := false
		for i := range numbers {
			if Validline(numbers) || Validline(RemoveIndex(numbers, i)) {
				valid = true
			}
		}
		if valid {
			validLine += 1
		}
	}
	fmt.Println(validLine)
	return validLine
}

func Validline(line []int) bool {
	valid := true
	for i, num := range line {
		if i >= len(line)-1 {
			break
		}
		num2 := line[i+1]
		dist := int(math.Abs(float64(num - num2)))
		// check if difference is > 1 < 3
		if dist < 1 || dist > 3 {
			valid = false
			break
		}
		// check if always increasing or decreasing
		if (num-num2 < 0 && line[0]-line[1] > 0) || (num-num2 > 0 && line[0]-line[1] < 0) {
			valid = false
			break
		}
	}
	return valid
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
