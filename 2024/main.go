package main

import ( // "context"
	day1 "adventofcode/day1"
	day2 "adventofcode/day2"
	day3 "adventofcode/day3"
	"fmt"
)

type AdventOfCode struct{}

// Executes the puzzle for all days
func main() {
	fmt.Println("day 1 part 1")
	fmt.Println(day1.Part1("1_1.txt"))
	fmt.Println("day 1 part 2")
	fmt.Println(day1.Part2("1_1.txt"))

	fmt.Println("day 2 part 1")
	fmt.Println(day2.Part1("2_1.txt"))
	fmt.Println("day 2 part 2")
	fmt.Println(day2.Part2("2_1.txt"))

	fmt.Println("day 3 part 1")
	fmt.Println(day3.Part1("3_1.txt"))
	fmt.Println("day 3 part 2")
	fmt.Println(day3.Part2("3_1.txt"))
}
