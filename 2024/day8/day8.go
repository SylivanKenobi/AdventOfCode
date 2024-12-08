package day7

import (
	"adventofcode/utils"
	"regexp"
	"strconv"
	"strings"
)

func Part1(file string) int {
	lines := utils.ReadFileByLine(file)
	antennas := Antennas(lines)

	antinodes := make(map[string]int)
	for _, antenna := range antennas {
		for i, point := range antenna {
			for j, pointj := range antenna {
				if i == j {
					continue
				}
				pathx := point[0] - pointj[0]
				pathy := point[1] - pointj[1]
				anti1 := []int{point[0] - pathx, point[1] - pathy}
				anti2 := []int{point[0] + pathx, point[1] + pathy}

				if validAntiPart1(anti1[0], anti1[1], len(lines[0]), len(lines), point[0], point[1], pointj[0], pointj[1]) {
					antinodes[strings.Join([]string{strconv.Itoa(anti1[0]), strconv.Itoa(anti1[1])}, "|")] = 1
				}
				if validAntiPart1(anti2[0], anti2[1], len(lines[0]), len(lines), point[0], point[1], pointj[0], pointj[1]) {
					antinodes[strings.Join([]string{strconv.Itoa(anti2[0]), strconv.Itoa(anti2[1])}, "|")] = 1
				}
			}
		}
	}
	return len(antinodes)
}

func Part2(file string) int {
	lines := utils.ReadFileByLine(file)
	antennas := Antennas(lines)
	antinodes := make(map[string]int)
	for _, antenna := range antennas {
		for i, point := range antenna {
			for j, pointj := range antenna {
				if i == j {
					continue
				}
				pathx := point[0] - pointj[0]
				pathy := point[1] - pointj[1]
				anti1 := []int{point[0] - pathx, point[1] - pathy}
				anti2 := []int{point[0] + pathx, point[1] + pathy}
				oneOut := false
				twoOut := false
				for !oneOut || !twoOut {
					if validAntiPart2(anti1[0], anti1[1], len(lines[0]), len(lines)) {
						antinodes[strings.Join([]string{strconv.Itoa(anti1[0]), strconv.Itoa(anti1[1])}, "|")] = 1
					} else {
						oneOut = true
					}

					if validAntiPart2(anti2[0], anti2[1], len(lines[0]), len(lines)) {
						antinodes[strings.Join([]string{strconv.Itoa(anti2[0]), strconv.Itoa(anti2[1])}, "|")] = 1
					} else {
						twoOut = true
					}
					anti1 = []int{anti1[0] - pathx, anti1[1] - pathy}
					anti2 = []int{anti2[0] + pathx, anti2[1] + pathy}
				}
			}
		}
	}
	return len(antinodes)
}

func Antennas(lines []string) map[string][][]int {
	antennas := make(map[string][][]int)
	antRe := regexp.MustCompile(`[0-9a-zA-Z]`)
	for i, line := range lines {
		antennaI := antRe.FindAllStringIndex(line, -1)
		antenna := antRe.FindAllString(line, -1)
		for j, a := range antenna {
			val, _ := antennas[a]
			xy := []int{antennaI[j][0], i}
			val = append(val, xy)
			antennas[a] = val
		}
	}
	return antennas
}

func validAntiPart1(x int, y int, xlimit int, ylimit int, antx1 int, anty1 int, antx2 int, anty2 int) bool {
	return (x >= 0 && x < xlimit && y >= 0 && y < ylimit && x != antx1 && x != antx2 && y != anty1 && y != anty2)
}

func validAntiPart2(x int, y int, xlimit int, ylimit int) bool {
	return (x >= 0 && x < xlimit && y >= 0 && y < ylimit)
}
