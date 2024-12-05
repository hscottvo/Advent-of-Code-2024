package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkLine(line []string) bool {
	curr, err := strconv.Atoi(line[0])
	if err != nil {
		panic(err)
	}
	next, err := strconv.Atoi(line[1])
	if err != nil {
		panic(err)
	}

	var increasing int

	if curr < next {
		increasing = 1

	} else if curr > next {
		increasing = -1
	} else {
		return false
	}

	for i := range len(line) - 1 {
		currNum, err := strconv.Atoi(line[i])
		if err != nil {
			panic(err)
		}

		nextNum, err := strconv.Atoi(line[i+1])
		if err != nil {
			panic(err)
		}

		nDiff := (nextNum - currNum) * increasing

		if nDiff < 1 || nDiff > 3 {
			return false
		}

	}

	return true
}

func isSafe(a int, b int, increasing int) bool {

	diff := (b - a) * increasing

	return diff > 0 && diff < 4
}

func part1() {
	data, err := os.ReadFile("./1.txt")
	if err != nil {
		panic(err)
	}

	str := string(data)
	lines := strings.Split(strings.TrimSpace(str), "\n")
	count := 0
	for _, i := range lines {
		if checkLine(strings.Split(i, " ")) {
			count++

		}
	}
	fmt.Println(count)
}

func checkLine2(line []string, start int, skip2 bool) bool {
	dampened := start != 0 || skip2

	first, err := strconv.Atoi(line[0])
	if err != nil {
		panic(err)
	}

	nextIdx := 1
	if skip2 {
		nextIdx = 2
	}
	second, err := strconv.Atoi(line[nextIdx])

	var increasing int
	if first < second {
		increasing = 1
	} else if first > second {
		increasing = -1
	} else {
		return false
	}

	if !isSafe(first, second, increasing) {
		return false
	}

	i := nextIdx
	for i < len(line)-1 {
		curr, err := strconv.Atoi(line[i])
		if err != nil {
			panic(err)
		}
		next, err := strconv.Atoi(line[i+1])
		if err != nil {
			panic(err)
		}

		if isSafe(curr, next, increasing) {
			i++
			continue
		}

		// if unsafe and dampened
		if dampened {
			return false
		}

		// if unsafe, undampened, and can dampen the last one
		if i+2 == len(line) {
			fmt.Printf("\tSkipping %s at index %d\n", line[i+1], i+1)
			fmt.Println("\t", line)
			return true
		}

		// otherwise, try the dampened case
		next, err = strconv.Atoi(line[i+2])
		if err != nil {
			panic(err)
		}
		if isSafe(curr, next, increasing) {
			fmt.Printf("Skipping %s at index %d\n", line[i+1], i+1)
			fmt.Println(line)
			dampened = true
			i = i + 2
			continue
		}
		return false

	}

	// if dampened {
	// 	fmt.Println(line)
	// }
	return true
}

func part2() {
	data, err := os.ReadFile("./1.txt")
	if err != nil {
		panic(err)
	}
	str := string(data)
	lines := strings.Split(strings.TrimSpace(str), "\n")
	count := 0
	for _, i := range lines {
		if checkLine2(strings.Split(i, " "), 0, false) {
			count++

		} else if checkLine2(strings.Split(i, " "), 1, false) {
			count++
		} else if checkLine2(strings.Split(i, " "), 0, true) {
			count++
		}

	}
	fmt.Println(count)
	//
	test := "1 1 2 3"
	fmt.Println(checkLine2(strings.Split(test, " "), 0, false))
	fmt.Println(checkLine2(strings.Split(test, " "), 1, false))
	fmt.Println(checkLine2(strings.Split(test, " "), 0, true))
}

func main() {
	part1()
	part2()

}
