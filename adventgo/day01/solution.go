package day01

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Solve1(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var l1, l2 []int
	for scanner.Scan() {
		row := scanner.Text()
		splits := strings.Fields(row)
		val1, _ := strconv.Atoi(splits[0])
		val2, _ := strconv.Atoi(splits[1])
		l1 = append(l1, val1)
		l2 = append(l2, val2)
	}

	// sort the slices
	sort.Ints(l1)
	sort.Ints(l2)

	total := 0
	for i := range l1 {
		total += absInt(l1[i] - l2[i])
	}

	fmt.Println("Answer:", total)
}

func Solve2(inputPath string) {
	fmt.Println("solve part 2 of day 1")
}
