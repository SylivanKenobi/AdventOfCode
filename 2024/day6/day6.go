package day6

import (
	"adventofcode/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Always turn right till stopped
func Part1(file string) []string {
	lines := utils.ReadFileByLine(file)

	upRe := regexp.MustCompile(`\^`)
	dwRe := regexp.MustCompile(`v`)
	riRe := regexp.MustCompile(`>`)
	leRe := regexp.MustCompile(`<`)
	// x,y,direction
	pos := make([]int, 3)
	for i, line := range lines {
		up := upRe.FindStringIndex(line)
		dw := dwRe.FindStringIndex(line)
		ri := riRe.FindStringIndex(line)
		le := leRe.FindStringIndex(line)
		if len(up) > 0 {
			pos[0], pos[1], pos[2] = up[0], i, 0
			break
		} else if len(dw) > 0 {
			pos[0], pos[1], pos[2] = dw[0], i, 2
			break
		} else if len(ri) > 0 {
			pos[0], pos[1], pos[2] = ri[0], i, 1
			break
		} else if len(le) > 0 {
			pos[0], pos[1], pos[2] = le[0], i, 3
			break
		}
	}
	var inc int
	var steps []string
	x := pos[0]
	y := pos[1]
	outOffBound := false
	for !outOffBound {
		if pos[2] == 1 || pos[2] == 2 {
			inc = 1
		} else {
			inc = -1
		}
		if string(lines[y][x]) == "#" {
			if pos[2] == 0 || pos[2] == 2 {
				y = y - inc
			} else {
				x = x - inc
			}
			if pos[2] == 3 {
				pos[2] = 0
			} else {
				pos[2] = pos[2] + 1
			}
			if pos[2] == 1 || pos[2] == 2 {
				inc = 1
			} else {
				inc = -1
			}
		} else {
			step := strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y)}, "|")
			steps = append(steps, step)
		}
		if pos[2] == 0 || pos[2] == 2 {
			y = y + inc
			outOffBound = (y < 0 || y > len(lines)-1)
		} else {
			x = x + inc
			outOffBound = (x < 0 || x > len(lines[0]))
		}
	}
	keys := make(map[string]bool)
	uniqueList := []string{}

	for _, entry := range steps {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueList = append(uniqueList, entry)
		}
	}
	return uniqueList
}

func Part2(file string, part1 []string) int {
	ogLines := utils.ReadFileByLine(file)

	result := 0

	upRe := regexp.MustCompile(`\^`)
	dwRe := regexp.MustCompile(`v`)
	riRe := regexp.MustCompile(`>`)
	leRe := regexp.MustCompile(`<`)
	// x,y,direction
	pos := make([]int, 3)
	for i, line := range ogLines {
		up := upRe.FindStringIndex(line)
		dw := dwRe.FindStringIndex(line)
		ri := riRe.FindStringIndex(line)
		le := leRe.FindStringIndex(line)
		if len(up) > 0 {
			pos[0], pos[1], pos[2] = up[0], i, 0
			break
		} else if len(dw) > 0 {
			pos[0], pos[1], pos[2] = dw[0], i, 2
			break
		} else if len(ri) > 0 {
			pos[0], pos[1], pos[2] = ri[0], i, 1
			break
		} else if len(le) > 0 {
			pos[0], pos[1], pos[2] = le[0], i, 3
			break
		}
	}

	var altLines [][]string
	for _, step := range part1 {
		tmpLines := make([]string, len(ogLines))
		copy(tmpLines, ogLines)
		xy := utils.StrToIntArr(strings.Split(step, "|"))
		bla := string(ogLines[xy[1]])
		out := []rune(bla)
		out[xy[0]] = rune('#')
		tmpLines[xy[1]] = string(out)
		altLines = append(altLines, tmpLines)
	}

	var inc int
	for _, lines := range altLines {
		x := pos[0]
		y := pos[1]
		dir := pos[2]
		outOffBound := false
		j := 0
		var steps []string
		for !outOffBound {
			if j > 0 && slices.Contains(steps, strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(dir)}, "|")) {
				result += 1
				break
			}
			if dir == 1 || dir == 2 {
				inc = 1
			} else {
				inc = -1
			}
			if string(lines[y][x]) == "#" {
				if dir == 0 || dir == 2 {
					y = y - inc
				} else {
					x = x - inc
				}
				if dir == 3 {
					dir = 0
				} else {
					dir = dir + 1
				}
				if dir == 1 || dir == 2 {
					inc = 1
				} else {
					inc = -1
				}

			} else {
				step := strings.Join([]string{strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(dir)}, "|")
				steps = append(steps, step)
			}
			if dir == 0 || dir == 2 {
				y = y + inc
				outOffBound = (y < 0 || y > len(lines)-1)
			} else {
				x = x + inc
				outOffBound = (x < 0 || x > len(lines[0])-1)
			}
			j++

		}
	}
	// 2093 too high
	return result
}
