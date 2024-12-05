package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	data, err := os.ReadFile("./1.txt")
	if err != nil {
		panic(err)
	}
	str := string(data)
	lines := strings.Split(strings.TrimSpace(str), "\n")

	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for idx, line := range lines {
		split := strings.Fields(line)
		left[idx], err = strconv.Atoi(split[0])
		right[idx], err = strconv.Atoi(split[1])

	}

	sort.Ints(left)
	sort.Ints(right)

	ret := 0

	for i := range len(left) {
		ret += int(math.Abs(float64(left[i]) - float64(right[i])))
	}
	fmt.Println(ret)

}

func part2() {
	data, err := os.ReadFile("./1.txt")
	if err != nil {
		panic(err)
	}
	str := string(data)
	lines := strings.Split(strings.TrimSpace(str), "\n")

	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for idx, line := range lines {
		split := strings.Fields(line)
		left[idx], err = strconv.Atoi(split[0])
		right[idx], err = strconv.Atoi(split[1])

	}

	m := make(map[int]int)

	for _, i := range right {
		m[i] += 1
	}

	ret := 0

	for _, i := range left {
		ret += m[i] * i
	}
	fmt.Println(ret)
}

func main() {
	// part1()
	part2()
}
