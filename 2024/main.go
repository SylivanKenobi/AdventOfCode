package main

import ( // "context"
	day1 "adventofcode/day1"
	day2 "adventofcode/day2"
	"fmt"
)

type AdventOfCode struct{}

// Executes the puzzle for all days
func main() {
	fmt.Println("day 1 part 1")
	day1.Part1("1_1.txt")
	fmt.Println("day 1 part 2")
	day1.Part2("1_1.txt")

	fmt.Println("day 1 part 1")
	day2.Part1("2_1.txt")
	fmt.Println("day 1 part 2")
	day2.Part2("2_1.txt")
}
