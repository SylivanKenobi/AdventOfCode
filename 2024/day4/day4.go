package day4

import (
	"adventofcode/utils"
	"strings"
)

func Part1(file string) int {
	lines := utils.ReadFileByLine(file)

	result := 0
	for i, line := range lines {
		for j, _ := range line {
			letters := strings.Split(line, "")
			if letters[j] == "X" || letters[j] == "S" {
				// right
				if len(letters)-3 > j {
					straight := strings.Join(letters[j:j+4], "")
					if Holliday(straight) {
						result += 1
					}
				}

				// down
				if len(lines)-3 > i {
					down := strings.Join([]string{string(lines[i][j]), string(lines[i+1][j]), string(lines[i+2][j]), string(lines[i+3][j])}, "")
					if Holliday(down) {
						// fmt.Println(down)
						result += 1
					}
				}

				// diagonal down right
				if len(lines)-3 > i && len(letters)-3 > j {
					ddr := strings.Join([]string{string(lines[i][j]), string(lines[i+1][j+1]), string(lines[i+2][j+2]), string(lines[i+3][j+3])}, "")
					if Holliday(ddr) {
						result += 1
					}
				}
				// diagonal down left
				if len(lines)-3 > i && 3 <= j {
					ddl := strings.Join([]string{string(lines[i][j]), string(lines[i+1][j-1]), string(lines[i+2][j-2]), string(lines[i+3][j-3])}, "")
					if Holliday(ddl) {
						result += 1
					}
				}
			}
		}
	}
	return result
}

func Part2(file string) int {
	// lines := utils.ReadFileByLine(file)
	result := 0

	return result
}

func Holliday(holliday string) bool {
	return (holliday == "XMAS" || holliday == "SAMX")
}
