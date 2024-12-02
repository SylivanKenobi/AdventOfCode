package day1

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1(file string) int {
	lines := utils.ReadFileByLine(file)
	var leftList []int
	var rightList []int
	for _, line := range lines {
		sides := strings.Split(line, "   ")
		left, _ := strconv.Atoi(sides[0])
		leftList = append(leftList, left)
		right, _ := strconv.Atoi(sides[1])
		rightList = append(rightList, right)
	}
	sort.Sort(sort.IntSlice(rightList))
	sort.Sort(sort.IntSlice(leftList))

	totalDist := 0
	for i, _ := range leftList {
		dist := leftList[i] - rightList[i]
		totalDist = int(math.Abs(float64(dist))) + totalDist
	}
	fmt.Println(totalDist)
	return totalDist
}

func Part2(file string) int {
	lines := utils.ReadFileByLine(file)
	var leftList []int
	var rightList []int
	for _, line := range lines {
		sides := strings.Split(line, "   ")
		left, _ := strconv.Atoi(sides[0])
		leftList = append(leftList, left)
		right, _ := strconv.Atoi(sides[1])
		rightList = append(rightList, right)
	}

	similarity := 0
	for _, left := range leftList {
		matches := 0
		for _, right := range rightList {
			if left == right {
				matches += 1
			}
		}
		similarity += (left * matches)
	}
	fmt.Println(similarity)
	return similarity
}
