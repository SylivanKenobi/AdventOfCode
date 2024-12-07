package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadFileByLine(filename string) []string {
	var file []string
	// open file
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		file = append(file, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return file
}

func StrToIntArr(input []string) []int {
	var r []int
	for _, n := range input {
		a, _ := strconv.Atoi(n)
		r = append(r, a)
	}
	return r
}
