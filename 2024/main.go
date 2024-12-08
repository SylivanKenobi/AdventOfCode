package main

import ( // "context"
	day1 "adventofcode/day1"
	day2 "adventofcode/day2"
	day3 "adventofcode/day3"
	day4 "adventofcode/day4"
	day5 "adventofcode/day5"
	day6 "adventofcode/day6"
	day7 "adventofcode/day7"
	day8 "adventofcode/day8"
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

	fmt.Println("day 4 part 1")
	fmt.Println(day4.Part1("4_1.txt"))
	fmt.Println("day 4 part 2")
	fmt.Println(day4.Part2("4_1.txt"))

	fmt.Println("day 5 part 1")
	fmt.Println(day5.Part1("5_1.txt"))
	fmt.Println("day 5 part 2")
	fmt.Println(day5.Part2("5_1.txt"))

	fmt.Println("day 6 part 1")
	sol := day6.Part1("6_1.txt")
	fmt.Println(len(sol))
	fmt.Println("day 6 part 2")
	fmt.Println(day6.Part2("6_1.txt", sol))

	fmt.Println("day 7 part 1")
	fmt.Println(day7.Part1("7_1.txt"))
	fmt.Println("day 7 part 2")
	fmt.Println(day7.Part2("7_1.txt"))

	fmt.Println("day 8 part 1")
	fmt.Println(day8.Part1("8_1.txt"))
	fmt.Println("day 8 part 2")
	fmt.Println(day8.Part2("8_1.txt"))
}
