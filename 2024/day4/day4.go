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
					if Holiday(straight) {
						result += 1
					}
				}

				// down
				if len(lines)-3 > i {
					down := strings.Join([]string{string(lines[i][j]), string(lines[i+1][j]), string(lines[i+2][j]), string(lines[i+3][j])}, "")
					if Holiday(down) {
						// fmt.Println(down)
						result += 1
					}
				}

				// diagonal down right
				if len(lines)-3 > i && len(letters)-3 > j {
					ddr := strings.Join([]string{string(lines[i][j]), string(lines[i+1][j+1]), string(lines[i+2][j+2]), string(lines[i+3][j+3])}, "")
					if Holiday(ddr) {
						result += 1
					}
				}
				// diagonal down left
				if len(lines)-3 > i && 3 <= j {
					ddl := strings.Join([]string{string(lines[i][j]), string(lines[i+1][j-1]), string(lines[i+2][j-2]), string(lines[i+3][j-3])}, "")
					if Holiday(ddl) {
						result += 1
					}
				}
			}
		}
	}
	return result
}

func Part2(file string) int {
	lines := utils.ReadFileByLine(file)

	result := 0
	for i, line := range lines {
		for j := range line {
			letters := strings.Split(line, "")
			if letters[j] == "M" || letters[j] == "S" {
				down := false
				up := false
				// diagonal down right
				if len(lines)-2 > i && len(letters)-2 > j {
					ddr := strings.Join([]string{string(lines[i][j]), string(lines[i+1][j+1]), string(lines[i+2][j+2])}, "")
					down = Mas(ddr)
				}
				// diagonal down left
				if len(lines)-2 > i && len(letters)-2 > j && down {
					ddl := strings.Join([]string{string(lines[i][j+2]), string(lines[i+1][j+1]), string(lines[i+2][j])}, "")
					up = Mas(ddl)
				}
				if up && down {
					result += 1
				}
			}
		}
	}
	return result
}

func Holiday(holliday string) bool {
	return (holliday == "XMAS" || holliday == "SAMX")
}

func Mas(holliday string) bool {
	return (holliday == "MAS" || holliday == "SAM")
}
